// Copyright 2024 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command_test

import (
	"testing"

	"github.com/kanisterio/safecli/command"
	"github.com/kanisterio/safecli/test"
	"gopkg.in/check.v1"
)

func TestOptionArgs(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&test.ArgumentSuite{Cmd: "cmd", Arguments: []test.ArgumentTest{
	{
		Name:        "NewOptionWithArgument",
		Argument:    command.NewOptionWithArgument("--option", "optArg1"),
		ExpectedCLI: []string{"cmd", "--option=optArg1"},
	},
	{
		Name:        "NewOptionWithArgument with empty name",
		Argument:    command.NewOptionWithArgument("", "optArg1"),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithArgument with invalid name",
		Argument:    command.NewOptionWithArgument("option", "optArg1"),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithArgument with empty argument",
		Argument:    command.NewOptionWithArgument("--option", ""),
		ExpectedCLI: []string{"cmd"},
	},

	{
		Name:        "NewOptionWithRedactedArgument",
		Argument:    command.NewOptionWithRedactedArgument("--option", "optArg1"),
		ExpectedCLI: []string{"cmd", "--option=optArg1"},
		ExpectedLog: "cmd --option=<****>",
	},
	{
		Name:        "NewOptionWithRedactedArgument with empty name",
		Argument:    command.NewOptionWithRedactedArgument("", "optArg1"),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithRedactedArgument with invalid name",
		Argument:    command.NewOptionWithRedactedArgument("option", "optArg1"),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithRedactedArgument with empty argument",
		Argument:    command.NewOptionWithRedactedArgument("--option", ""),
		ExpectedCLI: []string{"cmd"},
	},
}})
