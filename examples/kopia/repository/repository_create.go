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

package repository

import (
	"github.com/kanisterio/safecli"
	"github.com/kanisterio/safecli/examples/kopia/args"
	"github.com/kanisterio/safecli/examples/kopia/internal"
	"github.com/kanisterio/safecli/examples/kopia/internal/opts"
)

// CreateArgs represents the arguments for the `kopia repository create` command.
type CreateArgs struct {
	args.Common

	Hostname string
	Username string
}

// Create creates a new safecli.Builder for the `kopia repository create` command.
func Create(args CreateArgs) (*safecli.Builder, error) {
	return internal.NewKopiaCommand(opts.Common(args.Common),
		cmdRepository, cmdCreate,
		optHostname(args.Hostname),
		optUsername(args.Username),
	)
}
