// Copyright 2024 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command_test

import (
	"testing"

	"gopkg.in/check.v1"

	"github.com/pkg/errors"

	"github.com/kanisterio/safecli"
	"github.com/kanisterio/safecli/command"
	"github.com/kanisterio/safecli/test"
)

var (
	ErrArgument = errors.New("arg error")
)

// MockArg is a mock implementation of the Applier interface.
type MockArg struct {
	name string
	err  error
}

func (m *MockArg) Apply(cli safecli.CommandAppender) error {
	cli.AppendLoggable(m.name)
	return m.err
}

func TestArguments(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&test.ArgumentSuite{Cmd: "cmd", Arguments: []test.ArgumentTest{
	{
		Name:        "NewErrorArgument without error",
		Argument:    command.NewErrorArgument(nil),
		ExpectedCLI: []string{"cmd"},
	},
	{
		Name:        "NewErrorArgument with error",
		Argument:    command.NewErrorArgument(ErrArgument),
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: ErrArgument,
	},
	{
		Name:        "NewArgument",
		Argument:    command.NewArgument("arg1"),
		ExpectedCLI: []string{"cmd", "arg1"},
	},
	{
		Name:        "NewArgument with empty name",
		Argument:    command.NewArgument(""),
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidArgumentName,
	},
	{
		Name:        "NewRedactedArgument",
		Argument:    command.NewRedactedArgument("arg1"),
		ExpectedCLI: []string{"cmd", "arg1"},
		ExpectedLog: "cmd <****>",
	},
	{
		Name:        "NewRedactedArgument with empty name",
		Argument:    command.NewRedactedArgument(""),
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidArgumentName,
	},
	{
		Name: "NewArguments",
		Argument: command.NewArguments(
			command.NewArgument("arg1"),
			nil, // should be skipped
			command.NewRedactedArgument("arg2"),
		),
		ExpectedCLI: []string{"cmd", "arg1", "arg2"},
		ExpectedLog: "cmd arg1 <****>",
	},
	{
		Name:        "NewArguments without args",
		ExpectedCLI: []string{"cmd"},
	},
}})
