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

package command

import (
	"strings"

	"github.com/kanisterio/safecli"
)

// option defines an option with a given name.
// option name, ie: --option, -o
type option string

// Apply appends the option to the command.
func (o option) Apply(cmd safecli.CommandAppender) error {
	cmd.AppendLoggable(string(o))
	return nil
}

// NewOption creates a new option with a given name and enabled state.
func NewOption(name string, isEnabled bool) Applier {
	if err := validateOptionName(name); err != nil {
		return err
	} else if !isEnabled {
		return noopArgument{}
	}
	return option(name)
}

// NewToggleOption creates a new toggle option with a given enabled and disabled option name.
func NewToggleOption(enabledOpt, disabledOpt string, isEnabled bool) Applier {
	optName := disabledOpt
	if isEnabled {
		optName = enabledOpt
	}
	return NewOption(optName, true)
}

// validateOptionName validates the option name.
// It returns an error applier if the option name is empty
// or does not start with a hyphen prefix.
func validateOptionName(name string) Applier {
	if !strings.HasPrefix(name, "-") {
		return NewErrorArgument(ErrInvalidOptionName)
	}
	return nil
}
