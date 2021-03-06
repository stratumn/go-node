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

// Package service defines a service that deals with exit signals.
package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	logger "github.com/stratumn/go-node/core/log"

	logging "github.com/ipfs/go-log"
)

var (
	// ErrNotManager is returned when the connected service is not a
	// manager.
	ErrNotManager = errors.New("connected service is not a manager")
)

// log is the logger for the service.
var log = logging.Logger("signal")

// Manager represents a service manager.
type Manager interface {
	StopAll()
}

// Service is the Signal Handler service.
type Service struct {
	config *Config
	mgr    Manager
}

// Config contains configuration options for the Signal Handler service.
type Config struct {
	// Manager is the name of the manager service.
	Manager string `toml:"manager" comment:"The name of the manager service."`

	// AllowForcedShutdown allows forces shutdown by sending a second signal.
	AllowForcedShutdown bool `toml:"allow_forced_shutdown" comment:"Allow forced shutdown by sending second signal."`
}

// ID returns the unique identifier of the service.
func (s *Service) ID() string {
	return "signal"
}

// Name returns the human friendly name of the service.
func (s *Service) Name() string {
	return "Signal Handler"
}

// Desc returns a description of what the service does.
func (s *Service) Desc() string {
	return "Handles exit signals."
}

// Config returns the current service configuration or creates one with
// good default values.
func (s *Service) Config() interface{} {
	if s.config != nil {
		return *s.config
	}

	return Config{
		Manager:             "manager",
		AllowForcedShutdown: true,
	}
}

// SetConfig configures the service handler.
func (s *Service) SetConfig(config interface{}) error {
	conf := config.(Config)
	s.config = &conf
	return nil
}

// Needs returns the set of services this service depends on.
func (s *Service) Needs() map[string]struct{} {
	needs := map[string]struct{}{}
	needs[s.config.Manager] = struct{}{}

	return needs
}

// Plug sets the connected services.
func (s *Service) Plug(handlers map[string]interface{}) error {
	var ok bool

	if s.mgr, ok = handlers[s.config.Manager].(Manager); !ok {
		return errors.Wrap(ErrNotManager, s.config.Manager)
	}

	return nil
}

// Run starts the service.
func (s *Service) Run(ctx context.Context, running, stopping func()) error {
	sigint := make(chan os.Signal, 2)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

	running()

	select {
	case sig := <-sigint:
		log.Event(ctx, "exitSignal", logging.Metadata{"signal": sig})

		// Without a goroutine we would have a race condition because
		// the manager wouldn't be able to shut the service down.
		go s.exit(ctx)
		<-ctx.Done()

	case <-ctx.Done():
		stopping()
		signal.Stop(sigint)
	}

	return errors.WithStack(ctx.Err())
}

// exit handles exiting the process.
func (s *Service) exit(ctx context.Context) {
	log.Event(ctx, "exit")

	go func() {
		s.mgr.StopAll()
		logger.Close()
		fmt.Println("Goodbye.")
		os.Exit(0)
	}()

	if !s.config.AllowForcedShutdown {
		select {}
	}

	sigc := make(chan os.Signal, 2)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	fmt.Println("\nShutting down gracefully...\nPress Ctrl^C to force shutdown.")

	<-sigc
	fmt.Println("Forcing shutdown.")
	os.Exit(1)
}
