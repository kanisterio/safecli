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

	"github.com/kanisterio/safecli/command"
	"github.com/kanisterio/safecli/test"
	"gopkg.in/check.v1"
)

func TestOptions(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&test.ArgumentSuite{Cmd: "cmd", Arguments: []test.ArgumentTest{
	{
		Name:        "NewOption",
		Argument:    command.NewOption("--option", true),
		ExpectedCLI: []string{"cmd", "--option"},
	},
	{
		Name:        "NewOption disabled",
		Argument:    command.NewOption("--option", false),
		ExpectedCLI: []string{"cmd"},
	},
	{
		Name:        "NewOption with empty name",
		Argument:    command.NewOption("", false),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOption with invalid name",
		Argument:    command.NewOption("arg1", true),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewToggleOption",
		Argument:    command.NewToggleOption("--option", "--no-option", true),
		ExpectedCLI: []string{"cmd", "--option"},
	},
	{
		Name:        "NewToggleOption with empty name",
		Argument:    command.NewToggleOption("", "--no-option", true),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewToggleOption with invalid name",
		Argument:    command.NewToggleOption("option", "--no-option", true),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewToggleOption",
		Argument:    command.NewToggleOption("--option", "--no-option", false),
		ExpectedCLI: []string{"cmd", "--no-option"},
	},
	{
		Name:        "NewToggleOption with empty name",
		Argument:    command.NewToggleOption("--option", "", false),
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewToggleOption with invalid name",
		Argument:    command.NewToggleOption("--option", "no-option", false),
		ExpectedErr: command.ErrInvalidOptionName,
	},
}})
