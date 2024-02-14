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

import "github.com/kanisterio/safecli"

// Apply defines the interface for applying arguments to the command.
type Applier interface {
	// Apply applies arguments to the command.
	Apply(safecli.CommandAppender) error
}

// Apply appends multiple arguments to the command.
// If any of the arguments encounter an error during the Apply process,
// the error is returned and no changes are made to the command.
// If no error, the arguments are appended to the command.
func Apply(cmd safecli.CommandAppender, args ...Applier) error {
	// create a new subcmd builder which will be used to apply the arguments
	// to avoid mutating the command if an error is encountered.
	subcmd := safecli.NewBuilder()
	for _, arg := range args {
		if arg == nil { // if the param is nil, skip it
			continue
		}
		if err := arg.Apply(subcmd); err != nil {
			return err
		}
	}
	cmd.Append(subcmd)
	return nil
}
