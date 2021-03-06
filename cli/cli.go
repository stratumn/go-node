// Copyright © 2017-2018 Stratumn SAS
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

/*
Package cli defines types for Stratumn Node's command line interface.

It comes with only of handful of builtin commands. The bulk of commands are
reflected from the API.

The main type is the CLI struct, which wraps everything needed to run the
command line interface. It can, amongst other things, make suggestions for
auto-completion and connect to a Stratumn node.

The CLI needs a Console and a Prompt. The console is responsible for rendering
text. The Prompt is responsible for getting user input.

The Prompt is also responsible for calling the CLI's Exec method to execute
a command, and the Suggest method to make suggestions for auto-completion.

The Suggest method should be given a Content which allows it to read the
current text and returns a slice suggestions with the type Suggest.

BasicCmd is a type that allows creating simple commands that cover most use
cases.
*/
package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/stratumn/go-node/core/cfg"
	"github.com/stratumn/go-node/script"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
)

// StaticCmds is the list of builtin static CLI commands.
var StaticCmds = []Cmd{
	Addr,
	Bang,
	Connect,
	Disconnect,
	Echo,
	Exit,
	Help,
	Silent,
	Version,
}

// initScript is executed when the CLI is launched.
const initScript = `
echo
echo Stratumn Node CLI
echo
cli-version --git-commit-length 7
echo

; Only visible in debug mode.
echo --log debug Debug output is enabled.

; Connect to the API.
api-connect

echo
echo Enter "'help'" to list available commands.
echo Enter "'exit'" to quit the command line interface.
echo Use the tab key for auto-completion.
echo
`

// Available prompts.
var promptsMu sync.Mutex
var prompts = map[string]func(context.Context, CLI){}

// registerPrompt registers a prompt.
func registerPrompt(name string, run func(context.Context, CLI)) {
	promptsMu.Lock()
	prompts[name] = run
	promptsMu.Unlock()
}

// Resolver resolves a symbol to its name unless it begins with a dollar sign.
//
// This is convenient for shell scripting because you can do:
//
//	echo hello world
//
// Instead of:
//
//	echo "hello world"
//
// On the other hand you must prefix a symbol with a dollar sign to refer to
// it:
//
//	let str "hello world"
//	echo $str
func Resolver(c *script.Closure, sym script.SExp) (script.SExp, error) {
	name := sym.MustSymbolVal()

	if !strings.HasPrefix(name, "$") {
		return script.ResolveName(c, sym)
	}

	v, ok := c.Get(sym.MustSymbolVal()[1:])
	if !ok {
		// There is no need to wrap the error because the closure will
		// not return the error and look for the value in the closure
		// instead.
		return nil, script.ErrSymNotFound
	}

	return v, nil
}

// cli implements the command line interface.
type cli struct {
	conf Config

	cons   *Console
	prompt func(context.Context, CLI)

	itr       *script.Interpreter
	reflector ServerReflector

	staticCmds []Cmd
	apiCmds    []Cmd
	allCmds    []Cmd

	eventListener     EventListener
	eventListenerCh   chan struct{}
	stopEventListener context.CancelFunc

	addr string
	conn *grpc.ClientConn

	// Hack to hide suggestions after executing a command.
	executed bool
}

// New create a new command line interface.
func New(configSet cfg.ConfigSet) (CLI, error) {
	config, ok := configSet["cli"].(Config)
	if !ok {
		return nil, errors.WithStack(ErrInvalidConfig)
	}

	cons := NewConsole(os.Stdout, config.EnableColorOutput)

	prompt, ok := prompts[config.PromptBackend]
	if !ok {
		return nil, errors.WithStack(ErrPromptNotFound)
	}

	c := cli{
		conf:            config,
		cons:            cons,
		prompt:          prompt,
		reflector:       NewServerReflector(cons, 0),
		staticCmds:      StaticCmds,
		allCmds:         StaticCmds,
		addr:            config.APIAddress,
		eventListenerCh: make(chan struct{}),
	}

	// Create the top closure.
	closure := script.NewClosure(
		// Add env variables to the closure.
		script.ClosureOptEnv(os.Environ()),
		// Resolve symbols without a dollar sign to their names.
		script.ClosureOptResolver(Resolver),
	)

	c.itr = script.NewInterpreter(
		script.InterpreterOptClosure(closure),
		// Add all builtin script functions.
		script.InterpreterOptBuiltinLibs,
		// Print errors.
		script.InterpreterOptErrorHandler(c.printError),
		// Print evaluated instruction values.
		script.InterpreterOptValueHandler(c.printValue),
	)

	// Register static commands.
	sortCmds(c.staticCmds)
	c.addCmdsToInterpreter(StaticCmds)

	c.cons.SetDebug(config.EnableDebugOutput)

	return &c, nil
}

// Start starts the command line interface until the user kills it.
func (c *cli) Start(ctx context.Context) {
	defer func() {
		if err := c.Disconnect(); err != nil {
			if errors.Cause(err) != ErrDisconnected {
				c.cons.Debugf("Could not disconnect: %s.", err)
			}
		}
	}()

	c.Run(ctx, initScript)

	// Execute custom init scripts.
	for _, filename := range c.conf.InitScripts {
		c.runFile(ctx, filename)
	}

	c.prompt(ctx, c)
}

// Config returns the configuration.
func (c *cli) Config() Config {
	return c.conf
}

// Console returns the console.
func (c *cli) Console() *Console {
	return c.cons
}

// Commands returns all the commands.
func (c *cli) Commands() []Cmd {
	return c.allCmds
}

// Address returns the address of the API server.
func (c *cli) Address() string {
	return c.addr
}

// dialOpts builds the options to dial the API.
func (c *cli) dialOpts(ctx context.Context, address string) ([]grpc.DialOption, error) {
	opts := []grpc.DialOption{
		// Makes the call block till the connection is made.
		grpc.WithBlock(),
		// Use multiaddr dialer.
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			addr, err := ma.NewMultiaddr(address)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			dialer := manet.Dialer{}
			conn, err := dialer.DialContext(ctx, addr)
			return conn, errors.WithStack(err)
		}),
	}

	cert, override := c.conf.TLSCertificateFile, c.conf.TLSServerNameOverride
	if cert != "" {
		c.cons.Successln("TLS is enabled.")
		if override != "" {
			c.cons.Warningln("WARNING: API server authenticity is not checked because TLS server name override is enabled.")
		}

		creds, err := credentials.NewClientTLSFromFile(cert, override)
		if err != nil {
			return nil, err
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		c.cons.Warningln("WARNING: API connection is not encrypted because TLS is disabled.")
		opts = append(opts, grpc.WithInsecure())
	}

	return opts, nil
}

// Connect connects to the API server.
func (c *cli) Connect(ctx context.Context, addr string) error {
	if err := c.Disconnect(); err != nil && errors.Cause(err) != ErrDisconnected {
		return err
	}

	if addr != "" {
		c.addr = addr
	}
	if c.addr == "" {
		c.addr = c.conf.APIAddress
	}

	dto, err := time.ParseDuration(c.conf.DialTimeout)
	if err != nil {
		return errors.WithStack(err)
	}

	dialCtx, dialCancel := context.WithTimeout(ctx, dto)
	defer dialCancel()

	opts, err := c.dialOpts(dialCtx, c.addr)
	if err != nil {
		return err
	}

	conn, err := grpc.DialContext(dialCtx, c.addr, opts...)
	if err != nil {
		return errors.WithStack(err)
	}

	// Reflect API commands.
	c.apiCmds, err = c.reflector.Reflect(ctx, conn)
	if err != nil {
		if err := conn.Close(); err != nil {
			c.cons.Debugf("Could not close connection: %s.\n", err)
		}
		return err
	}

	c.eventListener = NewConsoleRPCEventListener(
		c.cons,
		conn,
	)

	// The input ctx will be cancelled once the connection is established.
	// We need a longer context for the event listeners.
	eventCtx, cancel := context.WithCancel(context.Background())
	c.stopEventListener = cancel
	go func(el EventListener) {
		err := el.Start(eventCtx)
		if err != nil {
			c.cons.Debugf("Event listener error: %s.\n", err.Error())
		}
		c.eventListenerCh <- struct{}{}
	}(c.eventListener)

	c.allCmds = append(c.staticCmds, c.apiCmds...)
	sortCmds(c.allCmds)

	c.addCmdsToInterpreter(c.apiCmds)

	c.conn = conn

	return nil
}

// Disconnect closes the API client connection.
func (c *cli) Disconnect() error {
	if c.conn == nil {
		return errors.WithStack(ErrDisconnected)
	}

	c.removeCmdsFromInterpreter(c.apiCmds)

	if c.stopEventListener != nil {
		c.stopEventListener()
		<-c.eventListenerCh
	}

	conn := c.conn
	c.conn = nil
	c.allCmds = c.staticCmds
	c.apiCmds = nil
	c.eventListener = nil

	return errors.WithStack(conn.Close())
}

// printError prints the error if it isn't nil.
func (c *cli) printError(err error) {
	if err == nil {
		return
	}

	c.cons.Errorf("Error: %s.\n", err)

	if stack := StackTrace(err); len(stack) > 0 {
		c.cons.Debugf("%+v\n", stack)
	}

	if useError, ok := errors.Cause(err).(*UseError); ok {
		if use := useError.Use(); use != "" {
			c.cons.Print("\n" + use)
		}
	}
}

// printValue prints an interpreted value.
func (c *cli) printValue(val script.SExp) {
	if val == nil {
		return
	}

	// Handle strings differently so they are not surrounded by quotes.
	if val.UnderlyingType() == script.SExpString {
		str := val.MustStringVal()
		if str == "" {
			return
		}

		c.cons.Print(str)
		if !strings.HasSuffix(str, "\n") {
			c.cons.Println()
		}

		return
	}

	c.cons.Println(fmt.Sprint(val))
}

// Exec executes the given input.
func (c *cli) Exec(ctx context.Context, in string) error {
	defer func() { c.executed = true }()

	return c.itr.EvalInput(ctx, in)
}

// Run executes the given input, handling errors and cancellation signals.
func (c *cli) Run(ctx context.Context, in string) {
	defer func() { c.executed = true }()

	execCtx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

	done := make(chan struct{})
	sigc := make(chan os.Signal, 2)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Errors are already printed, so ignore this one.
		_ = c.Exec(execCtx, in)
		close(done)
	}()

	// Handle exit conditions.
	select {
	case <-sigc:
		cancel()
		<-done
	case <-done:
	}
}

// runFile runs a script from a file.
func (c *cli) runFile(ctx context.Context, filename string) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		c.printError(errors.Wrap(err, filename))
		return
	}

	c.Run(ctx, string(in))
}

// Suggest finds all command suggestions.
func (c *cli) Suggest(cnt Content) []Suggest {
	var sug []Suggest

	for _, v := range c.allCmds {
		sug = append(sug, v.Suggest(cnt)...)
	}

	// Sort by text.
	sort.Slice(sug, func(i, j int) bool {
		return sug[i].Text < sug[j].Text
	})

	return sug
}

// DidJustExecute returns true the first time it is called after a command
// executed. This is a hack used by the VT100 prompt to hide suggestions
// after a command was executed.
func (c *cli) DidJustExecute() bool {
	defer func() { c.executed = false }()

	return c.executed
}

// addCmdsToInterpreter adds commands to the interpreter.
func (c *cli) addCmdsToInterpreter(cmds []Cmd) {
	for _, cmd := range cmds {
		c.itr.AddFuncHandler(cmd.Name(), c.wrapCmdExec(cmd))
	}
}

// removeCmdsFromInterpreter removes commands from the interpreter.
func (c *cli) removeCmdsFromInterpreter(cmds []Cmd) {
	for _, cmd := range cmds {
		c.itr.DeleteFuncHandler(cmd.Name())
	}
}

// wrapCmdExec makes the exec function of a command compatible with the
// interpreter.
func (c *cli) wrapCmdExec(cmd Cmd) script.InterpreterFuncHandler {
	return func(ctx *script.InterpreterContext) (script.SExp, error) {
		val, err := cmd.Exec(ctx, c)

		if useError, ok := errors.Cause(err).(*UseError); ok {
			useError.use = cmd.LongUse()
		}

		return val, err
	}
}

// sortCmds sorts a slice of commands by name.
func sortCmds(cmds []Cmd) {
	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].Name() < cmds[j].Name()
	})
}
