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

// errorArgument is a simple implementation of the Applier interface
// that always returns an error when applied.
type errorArgument struct {
	err error // error to return when applied
}

// Apply does nothing except return an error if one is set.
func (e errorArgument) Apply(cmd safecli.CommandAppender) error {
	return e.err
}

// NewErrorArgument creates a new argument with a given error.
// It is useful for creating an argument that always fails when applied.
func NewErrorArgument(err error) Applier {
	return errorArgument{err: err}
}

// noopArgument is a simple implementation of the Applier interface that does nothing.
type noopArgument struct{}

func (noopArgument) Apply(safecli.CommandAppender) error {
	return nil
}

// NewNoopArgument creates a new argument that does nothing when applied.
func NewNoopArgument() Applier {
	return noopArgument{}
}

// argument defines an argument with the given name.
// If the argument is redacted, it is appended as redacted.
type argument struct {
	name       string
	isRedacted bool
}

// Apply appends the argument to the command.
func (a argument) Apply(cmd safecli.CommandAppender) error {
	append := cmd.AppendLoggable
	if a.isRedacted {
		append = cmd.AppendRedacted
	}
	append(a.name)
	return nil
}

// newArgument creates a new argument with a given name and .
func newArgument(name string, isRedacted bool) Applier {
	if name == "" {
		return NewErrorArgument(ErrInvalidArgumentName)
	}
	return argument{
		name:       name,
		isRedacted: isRedacted,
	}
}

// NewArgument creates a new argument with a given name.
func NewArgument(name string) Applier {
	return newArgument(name, false)
}

// NewRedactedArgument creates a new redacted argument with a given name.
func NewRedactedArgument(name string) Applier {
	return newArgument(name, true)
}

// Arguments defines a collection of command arguments.
type Arguments []Applier

// Apply applies the flags to the CLI.
func (args Arguments) Apply(cli safecli.CommandAppender) error {
	return apply(cli, args...)
}

// NewArguments creates a new collection of arguments.
func NewArguments(args ...Applier) Applier {
	return Arguments(args)
}
