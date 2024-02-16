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
	"strings"
	"testing"

	"github.com/kanisterio/safecli/command"
	"gopkg.in/check.v1"
)

func TestCommand(t *testing.T) { check.TestingT(t) }

type CommandSuite struct{}

var _ = check.Suite(&CommandSuite{})

func (s *CommandSuite) TestCommandNewOK(c *check.C) {
	cli := []string{
		"cmd",
		"--log-level=info",
		"--password=secret",
		"arg",
		"--dest=/tmp/dir",
		"--read-only",
	}
	log := []string{
		"cmd",
		"--log-level=info",
		"--password=<****>",
		"arg",
		"--dest=/tmp/dir",
		"--read-only",
	}
	cmd, err := command.New("cmd", []command.Applier{
		command.NewOptionWithArgument("--log-level", "info"),
		command.NewOptionWithRedactedArgument("--password", "secret"),
		command.NewArgument("arg"),
		command.NewOptionWithArgument("--dest", "/tmp/dir"),
		command.NewOption("--read-only", true),
	}...)
	c.Assert(err, check.IsNil)
	c.Assert(cmd.Build(), check.DeepEquals, cli)
	c.Assert(cmd.String(), check.Equals, strings.Join(log, " "))
}

func (s *CommandSuite) TestCommandNewError(c *check.C) {
	cmd, err := command.New("cmd", []command.Applier{
		command.NewOptionWithArgument("--log-level", "info"),
		command.NewOptionWithRedactedArgument("--password", "secret"),
		command.NewArgument(""), // error argument
	}...)
	c.Assert(cmd, check.IsNil)
	c.Assert(err, check.Equals, command.ErrInvalidArgumentName)
}
