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
	"github.com/kanisterio/safecli"
)

// optionArg defines an option with the argument.
type optionArg struct {
	name       string
	arg        string
	isRedacted bool
}

// Apply appends the option to the command.
// If the option argument is redacted, it is appended as redacted.
func (o optionArg) Apply(cmd safecli.CommandAppender) error {
	// append the option and the argument to the command
	if o.isRedacted {
		cmd.AppendRedactedKV(o.name, o.arg)
	} else {
		cmd.AppendLoggableKV(o.name, o.arg)
	}
	return nil
}

// newOptionArg creates a new option with a given option name and redacted/non-redacted argument.
func newOptionArg(name, arg string, isArgRedacted bool) Applier {
	if err := validateOptionName(name); err != nil {
		return err
	} else if arg == "" {
		return noopArgument{}
	}
	return optionArg{
		name:       name,
		arg:        arg,
		isRedacted: isArgRedacted,
	}
}

// NewOptionWithArgument creates a new option with a given option name and argument.
func NewOptionWithArgument(name, arg string) Applier {
	return newOptionArg(name, arg, false)
}

// NewOptionWithRedactedArgument creates a new option with a given option name and redacted argument.
func NewOptionWithRedactedArgument(name, arg string) Applier {
	return newOptionArg(name, arg, true)
}
