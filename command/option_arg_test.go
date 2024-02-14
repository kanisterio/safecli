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
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithArgument with invalid name",
		Argument:    command.NewOptionWithArgument("option", "optArg1"),
		ExpectedCLI: []string{"cmd"},
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
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithRedactedArgument with invalid name",
		Argument:    command.NewOptionWithRedactedArgument("option", "optArg1"),
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOptionWithRedactedArgument with empty argument",
		Argument:    command.NewOptionWithRedactedArgument("--option", ""),
		ExpectedCLI: []string{"cmd"},
	},
}})
