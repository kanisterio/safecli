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
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewOption with invalid name",
		Argument:    command.NewOption("arg1", true),
		ExpectedCLI: []string{"cmd"},
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
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewToggleOption with invalid name",
		Argument:    command.NewToggleOption("option", "--no-option", true),
		ExpectedCLI: []string{"cmd"},
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
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
	{
		Name:        "NewToggleOption with invalid name",
		Argument:    command.NewToggleOption("--option", "no-option", false),
		ExpectedCLI: []string{"cmd"},
		ExpectedErr: command.ErrInvalidOptionName,
	},
}})
