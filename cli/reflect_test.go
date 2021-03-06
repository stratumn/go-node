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

package cli

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/jhump/protoreflect/desc"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/stratumn/go-node/cli/grpc/test"
	"github.com/stratumn/go-node/core/netutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type reflectorTest struct {
	name        string
	r           Reflector
	supported   *desc.FieldDescriptor
	unsupported *desc.FieldDescriptor
	arg         []reflectorParseTest
	flag        []reflectorParseTest
	pretty      []reflectorPrettyTest
}

type reflectorParseTest struct {
	text  string
	want  interface{}
	fails bool
}

type reflectorPrettyTest struct {
	val   interface{}
	want  string
	fails bool
}

var testTime time.Time

func init() {
	v, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-12-13 17:22:06.403954641 +0100 CET")
	if err != nil {
		panic(err)
	}

	testTime = v.UTC()
}

func TestReflectors(t *testing.T) {
	msg, err := desc.LoadMessageDescriptorForMessage(&test.Message{})
	require.NoError(t, err, "failed to load message")

	tests := []reflectorTest{{
		"string",
		NewStringReflector(),
		msg.FindFieldByName("str"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"test", "test", false},
			{" test", "test", false},
		},
		[]reflectorParseTest{
			{"cmd --str test", "test", false},
		},
		[]reflectorPrettyTest{
			{"test", "test", false},
		},
	}, {
		"string repeated",
		NewStringReflector(),
		msg.FindFieldByName("str_repeated"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"a", []string{"a"}, false},
			{"a,b", []string{"a", "b"}, false},
			{"a, b", []string{"a", "b"}, false},
		},
		[]reflectorParseTest{
			{"cmd --str_repeated a --str_repeated b", []string{"a", "b"}, false},
		},
		[]reflectorPrettyTest{
			{[]string{"a", "b"}, "a,b", false},
		},
	}, {
		"bool",
		NewBoolReflector(),
		msg.FindFieldByName("boolean"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"true", true, false},
			{"True", true, false},
			{"false", false, false},
			{"fa", false, true},
		},
		[]reflectorParseTest{
			{"cmd --boolean true", true, false},
			{"cmd --boolean false", false, false},
			{"cmd --boolean bla", false, true},
			{"cmd", false, false},
		},
		[]reflectorPrettyTest{
			{true, "true", false},
			{false, "false", false},
		},
	}, {
		"bool repeated",
		NewBoolReflector(),
		msg.FindFieldByName("boolean_repeated"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"true,false", []bool{true, false}, false},
			{"true , false", []bool{true, false}, false},
			{"True", []bool{true}, false},
			{"fa", []bool{}, true},
		},
		[]reflectorParseTest{
			{"cmd --boolean_repeated true,false", []bool{true, false}, false},
			{"cmd --boolean_repeated true --boolean_repeated false", []bool{true, false}, false},
			{"cmd", []bool{}, false},
		},
		[]reflectorPrettyTest{
			{[]bool{true, false}, "true,false", false},
		},
	}, {
		"int32",
		NewInt32Reflector(),
		msg.FindFieldByName("i32"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"10", int32(10), false},
		},
		[]reflectorParseTest{
			{"cmd --i32 10", int32(10), false},
			{"cmd --i32 false", int32(0), true},
			{"cmd", int32(0), false},
		},
		[]reflectorPrettyTest{
			{int32(10), "10", false},
		},
	}, {
		"uint32",
		NewUint32Reflector(),
		msg.FindFieldByName("u32"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"10", uint32(10), false},
		},
		[]reflectorParseTest{
			{"cmd --u32 10", uint32(10), false},
			{"cmd --u32 false", uint32(0), true},
			{"cmd", uint32(0), false},
		},
		[]reflectorPrettyTest{
			{uint32(10), "10", false},
		},
	}, {
		"int64",
		NewInt64Reflector(),
		msg.FindFieldByName("i64"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"10", int64(10), false},
		},
		[]reflectorParseTest{
			{"cmd --i64 10", int64(10), false},
			{"cmd --i64 false", int64(0), true},
			{"cmd", int64(0), false},
		},
		[]reflectorPrettyTest{
			{int64(10), "10", false},
		},
	}, {
		"uint64",
		NewUint64Reflector(),
		msg.FindFieldByName("u64"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"10", uint64(10), false},
		},
		[]reflectorParseTest{
			{"cmd --u64 10", uint64(10), false},
			{"cmd --u64 false", uint64(0), true},
			{"cmd", uint64(0), false},
		},
		[]reflectorPrettyTest{
			{uint64(10), "10", false},
		},
	}, {
		"bytes",
		NewBytesReflector(),
		msg.FindFieldByName("buf"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"ff00", []byte{255, 0}, false},
			{"zz", []byte{}, true},
		},
		[]reflectorParseTest{
			{"cmd --buf ff00", []byte{255, 0}, false},
			{"cmd", []byte{}, false},
			{"cmd --buf zz", []byte{}, true},
		},
		[]reflectorPrettyTest{
			{[]byte{255, 0}, "ff00", false},
		},
	}, {
		"bytes repeated",
		NewBytesReflector(),
		msg.FindFieldByName("buf_repeated"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"ff00,ff", [][]byte{{255, 0}, {255}}, false},
			{"ff00, ff", [][]byte{{255, 0}, {255}}, false},
			{"zz", [][]byte{}, true},
		},
		[]reflectorParseTest{
			{"cmd --buf_repeated ff00 --buf_repeated ff", [][]byte{{255, 0}, {255}}, false},
			{"cmd", [][]byte{}, false},
			{"cmd --buf_repeated zz", [][]byte{}, true},
		},
		[]reflectorPrettyTest{
			{[][]byte{{255, 0}, {255}}, "ff00,ff", false},
		},
	}, {
		"enum",
		NewEnumReflector(),
		msg.FindFieldByName("enumeration"),
		msg.FindFieldByName("buf"),
		[]reflectorParseTest{
			{"A", int32(0), false},
			{"B", int32(1), false},
			{"1", int32(0), true},
		},
		[]reflectorParseTest{
			{"cmd --enumeration b", int32(1), false},
			{"cmd", int32(0), false},
			{"cmd --enumeration zz", int32(0), true},
		},
		[]reflectorPrettyTest{
			{int32(1), "B", false},
		},
	}, {
		"time",
		NewTimeReflector(),
		msg.FindFieldByName("time"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{testTime.String(), testTime.UnixNano(), false},
			{"1", int64(0), true},
		},
		[]reflectorParseTest{
			{"cmd --time \"" + testTime.String() + "\"", testTime.UnixNano(), false},
			{"cmd", int64(0), false},
			{"cmd --time zz", int64(0), true},
		},
		[]reflectorPrettyTest{
			{testTime.UnixNano(), testTime.String(), false},
		},
	}, {
		"time repeated",
		NewTimeReflector(),
		msg.FindFieldByName("time_repeated"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{fmt.Sprintf("%q,%q", testTime, testTime), []int64{
				testTime.UnixNano(),
				testTime.UnixNano(),
			}, false},
			{"a,b", []int64{}, true},
		},
		[]reflectorParseTest{
			{
				fmt.Sprintf(
					"cmd --time_repeated %q --time_repeated %q",
					testTime, testTime,
				),
				[]int64{testTime.UnixNano(), testTime.UnixNano()},
				false,
			},
			{"cmd", []int64{}, false},
			{"cmd --time_repeated zz", []int64{}, true},
		},
		[]reflectorPrettyTest{
			{
				[]int64{testTime.UnixNano(), testTime.UnixNano()},
				fmt.Sprintf("%s,%s", testTime, testTime),
				false,
			},
		},
	}, {
		"duration",
		NewDurationReflector(),
		msg.FindFieldByName("duration"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"1m", int64(time.Minute), false},
			{"1", int64(0), true},
		},
		[]reflectorParseTest{
			{"cmd --duration 1m", int64(time.Minute), false},
			{"cmd", int64(0), false},
			{"cmd --duration zz", int64(0), true},
		},
		[]reflectorPrettyTest{
			{int64(time.Minute), "1m0s", false},
		},
	}, {
		"duration repeated",
		NewDurationReflector(),
		msg.FindFieldByName("duration_repeated"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"1m,1s", []int64{int64(time.Minute), int64(time.Second)}, false},
			{"a,b", []int64{}, true},
		},
		[]reflectorParseTest{
			{
				"cmd --duration_repeated 1m --duration_repeated 1s",
				[]int64{int64(time.Minute), int64(time.Second)},
				false,
			},
			{"cmd", []int64{}, false},
			{"cmd --duration_repeated zz", []int64{}, true},
		},
		[]reflectorPrettyTest{
			{[]int64{int64(time.Minute), int64(time.Second)}, "1m0s,1s", false},
		},
	}, {
		"base58",
		NewBase58Reflector(),
		msg.FindFieldByName("base58"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"QmVhJVRSYHNSHgR9dJNbDxu6G7GPPqJAeiJoVRvcexGNf9", []byte{
				18, 32, 109, 76, 36, 229, 32, 176,
				53, 31, 169, 189, 190, 37, 119, 25,
				79, 187, 126, 16, 211, 82, 93, 216,
				194, 134, 220, 138, 66, 252, 106, 50,
				58, 192,
			}, false},
			{"!1", []byte{}, true},
		},
		[]reflectorParseTest{
			{"cmd --base58 QmVhJVRSYHNSHgR9dJNbDxu6G7GPPqJAeiJoVRvcexGNf9", []byte{
				18, 32, 109, 76, 36, 229, 32, 176,
				53, 31, 169, 189, 190, 37, 119, 25,
				79, 187, 126, 16, 211, 82, 93, 216,
				194, 134, 220, 138, 66, 252, 106, 50,
				58, 192,
			}, false},
			{"cmd", []byte{}, false},
			{"cmd --base58 !", []byte{}, true},
		},
		[]reflectorPrettyTest{
			{[]byte{
				18, 32, 109, 76, 36, 229, 32, 176,
				53, 31, 169, 189, 190, 37, 119, 25,
				79, 187, 126, 16, 211, 82, 93, 216,
				194, 134, 220, 138, 66, 252, 106, 50,
				58, 192,
			}, "QmVhJVRSYHNSHgR9dJNbDxu6G7GPPqJAeiJoVRvcexGNf9", false},
		},
	}, {
		"bytesize",
		NewBytesizeReflector(),
		msg.FindFieldByName("bytesize"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"1K", uint64(1024), false},
			{"1", uint64(0), true},
		},
		[]reflectorParseTest{
			{"cmd --bytesize 1K", uint64(1024), false},
			{"cmd", uint64(0), false},
			{"cmd --bytesize zz", uint64(0), true},
		},
		[]reflectorPrettyTest{
			{uint64(1024), "1K", false},
		},
	}, {
		"byterate",
		NewByterateReflector(),
		msg.FindFieldByName("byterate"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"1K/s", uint64(1024), false},
			{"1/s", uint64(0), true},
		},
		[]reflectorParseTest{
			{"cmd --byterate 1K/s", uint64(1024), false},
			{"cmd", uint64(0), false},
			{"cmd --byterate zz", uint64(0), true},
		},
		[]reflectorPrettyTest{
			{uint64(1024), "1K/s", false},
		},
	}, {
		"maddr",
		NewMaddrReflector(),
		msg.FindFieldByName("multiaddr"),
		msg.FindFieldByName("boolean"),
		[]reflectorParseTest{
			{"/ip4/127.0.0.1/tcp/80", []byte{
				4, 127, 0, 0, 1, 6, 0, 80,
			}, false},
			{"!1", []byte{}, true},
		},
		[]reflectorParseTest{
			{"cmd --multiaddr /ip4/127.0.0.1/tcp/80", []byte{
				4, 127, 0, 0, 1, 6, 0, 80,
			}, false},
			{"cmd", []byte{}, false},
			{"cmd --multiaddr !", []byte{}, true},
		},
		[]reflectorPrettyTest{
			{[]byte{
				4, 127, 0, 0, 1, 6, 0, 80,
			}, "/ip4/127.0.0.1/tcp/80", false},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testReflector(t, tt)
		})
	}
}

func testReflector(t *testing.T, test reflectorTest) {
	assert.True(t, test.r.Supports(test.supported), "test.r.Supports(test.supported)")
	assert.False(t, test.r.Supports(test.unsupported), "test.r.Supports(test.unsupported)")

	testReflectorArg(t, test)
	testReflectorFlag(t, test)
	testReflectorPretty(t, test)
}

func testReflectorArg(t *testing.T, test reflectorTest) {
	if argReflector, ok := test.r.(ArgReflector); ok {
		for _, argTest := range test.arg {
			v, err := argReflector.Parse(test.supported, argTest.text)
			if err != nil && !argTest.fails {
				assert.Failf(t, "unexpected error", "argReflector.Parse(test.supported, %q)", argTest.text)
			}

			assert.EqualValuesf(t, argTest.want, v, "argReflector.Parse(test.supported, %q)", argTest.text)
		}
	}
}

func testReflectorFlag(t *testing.T, test reflectorTest) {
	if flagReflector, ok := test.r.(FlagReflector); ok {
		for _, flagTest := range test.flag {
			r := csv.NewReader(strings.NewReader(flagTest.text))
			r.Comma = ' '
			args, err := r.Read()
			assert.NoError(t, err, "r.Read()")

			f := pflag.NewFlagSet(test.name, pflag.ContinueOnError)
			flagReflector.Flag(test.supported, f)
			f.Parse(args)

			v, err := flagReflector.ParseFlag(test.supported, f)
			if err != nil && !flagTest.fails {
				assert.Failf(t, "unexpected error", "flagReflector.ParseFlags(%q, f)", args)
			}

			assert.EqualValuesf(t, flagTest.want, v, "flagReflector.ParseFlags(%q, f)", args)
		}
	}
}

func testReflectorPretty(t *testing.T, test reflectorTest) {
	if resReflector, ok := test.r.(ResponseReflector); ok {
		for _, prettyTest := range test.pretty {
			s, err := resReflector.Pretty(test.supported, prettyTest.val)
			if err != nil && !prettyTest.fails {
				assert.Failf(t, "unexpected error", "resReflector.Pretty(test.supported, %v)", prettyTest.val)
			}

			assert.EqualValuesf(t, prettyTest.want, s, "resReflector.Pretty(test.supported, %v)", prettyTest.val)
		}
	}
}

type reflectServer struct{}

func (reflectServer) UnaryReq(ctx context.Context, req *test.Message) (*test.Message, error) {
	return req, nil
}

func (reflectServer) ServerStream(req *test.Message, ss test.Test_ServerStreamServer) error {
	return ss.Send(req)
}

func (reflectServer) NoExt(ctx context.Context, req *test.Message) (*test.Message, error) {
	return req, nil
}

func testReflectServer(ctx context.Context, t *testing.T, address string) error {
	gs := grpc.NewServer()
	reflection.Register(gs)

	test.RegisterTestServer(gs, reflectServer{})

	lis, err := netutil.Listen(address)
	if err != nil {
		return err
	}

	ch := make(chan error, 1)
	go func() {
		ch <- gs.Serve(lis)
	}()

	select {
	case err = <-ch:
	case <-ctx.Done():
		gs.GracefulStop()
	}

	if err != nil {
		err = errors.Cause(err)

		if e, ok := err.(*net.OpError); ok {
			if e.Op == "accept" {
				// Normal error.
				return nil
			}
		}
	}

	return err
}

type reflectorServerTest struct {
	name string
	cmd  string
	want string
	err  error
}

const reflectTestUsage = `
Usage:
  test-unaryreq <required field>

Flags:
      --base58 string                  base58 field
      --base58_repeated strings        base58 repeated field
      --boolean string                 bool field
      --boolean_repeated strings       bool repeated field
      --buf string                     bytes field
      --buf_repeated strings           bytes repeated field
      --byterate string                byterate field
      --byterate_repeated strings      byterate repeated field
      --bytesize string                bytesize field
      --bytesize_repeated strings      bytesize repeated field
      --duration string                duration field
      --duration_repeated strings      duration repeated field
      --enumeration string             enum field
      --enumeration_repeated strings   enum repeated field
      --field string                   Only display specified field
  -h, --help                           Invoke help on command
      --i32 string                     int32 field
      --i32_repeated strings           int32 repeated field
      --i64 string                     int64 field
      --i64_repeated strings           int64 repeated field
      --multiaddr string               multiaddr field
      --multiaddr_repeated strings     multiaddr repeated field
      --no-timeout                     Disable request timeout
      --noext string                   noext
      --str string                     string field
      --str_repeated strings           string repeated field
      --time string                    time field
      --time_repeated strings          time repeated field
      --u32 string                     uint32 field
      --u32_repeated strings           uint32 repeated field
      --u64 string                     uint64 field
      --u64_repeated strings           uint64 repeated field
`

var serverReflectorTests = []reflectorServerTest{{
	"unary",
	"test-unaryreq hello --boolean true --bytesize_repeated 1k --bytesize_repeated 1m",
	`
NOEXT
REQUIRED FIELD            hello
STRING FIELD
STRING REPEATED FIELD
BOOL FIELD                true
BOOL REPEATED FIELD
INT32 FIELD               0
INT32 REPEATED FIELD
UINT32 FIELD              0
UINT32 REPEATED FIELD
INT64 FIELD               0
INT64 REPEATED FIELD
UINT64 FIELD              0
UINT64 REPEATED FIELD
BYTES FIELD
BYTES REPEATED FIELD
ENUM FIELD                A
ENUM REPEATED FIELD
BASE58 FIELD
BASE58 REPEATED FIELD
MULTIADDR FIELD
MULTIADDR REPEATED FIELD
TIME FIELD                1970-01-01 00:00:00 +0000 UTC
TIME REPEATED FIELD
DURATION FIELD            0s
DURATION REPEATED FIELD
BYTESIZE FIELD            0
BYTESIZE REPEATED FIELD   1K,1M
BYTERATE FIELD            0/s
BYTERATE REPEATED FIELD
`,
	nil,
}, {
	"unary --field",
	"test-unaryreq hello --field req",
	"hello",
	nil,
}, {
	"stream",
	"test-serverstream hello",
	`
NOEXT | REQUIRED FIELD | STRING FIELD | STRING REPEATED FIELD | BOOL FIELD ...
------+----------------+--------------+-----------------------+------------+--
      | hello          |              |                       | false      ...
`,
	nil,
}, {
	"stream --stream",
	"test-serverstream --no-timeout --stream hello --boolean true --bytesize_repeated 1k --bytesize_repeated 1m",
	`
NOEXT
REQUIRED FIELD            hello
STRING FIELD
STRING REPEATED FIELD
BOOL FIELD                true
BOOL REPEATED FIELD
INT32 FIELD               0
INT32 REPEATED FIELD
UINT32 FIELD              0
UINT32 REPEATED FIELD
INT64 FIELD               0
INT64 REPEATED FIELD
UINT64 FIELD              0
UINT64 REPEATED FIELD
BYTES FIELD
BYTES REPEATED FIELD
ENUM FIELD                A
ENUM REPEATED FIELD
BASE58 FIELD
BASE58 REPEATED FIELD
MULTIADDR FIELD
MULTIADDR REPEATED FIELD
TIME FIELD                1970-01-01 00:00:00 +0000 UTC
TIME REPEATED FIELD
DURATION FIELD            0s
DURATION REPEATED FIELD
BYTESIZE FIELD            0
BYTESIZE REPEATED FIELD   1K,1M
BYTERATE FIELD            0/s
BYTERATE REPEATED FIELD
`,
	nil,
}, {
	"stream --field",
	"test-serverstream hello --field req",
	"hello",
	nil,
}, {
	"stream --stream --field",
	"test-serverstream hello --stream --field req",
	"hello",
	nil,
}, {
	"missing arg",
	"test-unaryreq",
	ansiRed + "Error: 1:1: test-unaryreq: invalid usage: missing argument(s).\n" + ansiReset + reflectTestUsage,
	errUse,
}, {
	"extra arg",
	"test-unaryreq a b",
	ansiRed + "Error: 1:1: test-unaryreq: invalid usage: unexpected argument(s): b.\n" + ansiReset + reflectTestUsage,
	errUse,
}, {
	"invalid flag",
	"test-unaryreq --boolean 1",
	ansiRed + "Error: 1:1: test-unaryreq: invalid usage: missing argument(s).\n" + ansiReset + reflectTestUsage,
	errUse,
}, {
	"invalid field",
	"test-unaryreq hello --field test",
	ansiRed + "Error: 1:1: test-unaryreq: test: the field was not found.\n" + ansiReset,
	ErrFieldNotFound,
}}

func TestServerReflector_Reflect(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	addr := testAddress()

	serverCh := make(chan error, 2)
	go func() {
		serverCh <- testReflectServer(ctx, t, addr)
	}()

	config := NewConfigurableSet().Configs()
	conf := config["cli"].(Config)
	conf.APIAddress = addr
	c, err := New(config)
	assert.NoError(t, err, "New(config)")

	// Set truncate width.
	c.(*cli).reflector = NewServerReflector(c.Console(), 78)

	c.Console().Writer = ioutil.Discard

	err = c.Connect(ctx, addr)
	assert.NoError(t, err, "c.Connect(ctx, addr)")

	for _, tt := range serverReflectorTests {
		t.Run(tt.name, func(t *testing.T) {
			testServerReflectorReflect(ctx, t, c, tt)
		})
	}

	cancel()

	err = <-serverCh
	assert.NoError(t, err, "testServer(ctx, t, addr)")
}

var (
	errAny = errors.New("any error")
	errUse = errors.New("usage error")
)

func testServerReflectorReflect(ctx context.Context, t *testing.T, c CLI, test reflectorServerTest) {
	buf := bytes.NewBuffer(nil)
	c.Console().Writer = buf

	err := errors.Cause(c.Exec(ctx, test.cmd))

	switch {
	case test.err == errAny && err != nil:
		// Pass.
	case test.err == errUse:
		_, ok := err.(*UseError)
		assert.True(t, ok, "err.(*UseError)")
	case err != test.err:
		assert.Failf(t, "unexpected error", "%s", test.cmd)
	}

	got, want := trimLines(buf.String()), trimLines(test.want)
	assert.Equal(t, want, got)
}

func trimLines(s string) string {
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}

	return strings.Join(lines, "\n")
}
