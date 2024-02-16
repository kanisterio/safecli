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
	"fmt"
	"time"

	"github.com/kanisterio/safecli/command"
	"github.com/kanisterio/safecli/examples/kopia/repository/storage/fs"
	"github.com/kanisterio/safecli/examples/kopia/repository/storage/s3"
)

var (
	cmdRepository = command.NewArgument("repository")

	subcmdCreate  = command.NewArgument("create")
	subcmdConnect = command.NewArgument("connect")
)

// optHostname creates a new option for the hostname of the repository.
func optHostname(h string) command.Applier {
	return command.NewOptionWithArgument("--override-hostname", h)
}

// optUsername creates a new option for the username of the repository.
func optUsername(u string) command.Applier {
	return command.NewOptionWithArgument("--override-username", u)
}

// optReadOnly creates a new option for the read-only mode of the repository.
func optReadOnly(readOnly bool) command.Applier {
	return command.NewOption("--read-only", readOnly)
}

// optPointInTime creates a new option for the point-in-time of the repository.
func optPointInTime(pit time.Time) command.Applier {
	if pit.IsZero() {
		return command.NewNoopArgument()
	}
	return command.NewOptionWithArgument("--point-in-time", pit.Format(time.RFC3339))
}

// optStorage creates a list of options for the specified storage location.
func optStorage(l Location) command.Applier {
	switch l.Provider {
	case ProviderFilesystem:
		return fs.New(l.MetaData)
	case ProviderS3:
		return s3.New(l.MetaData)
	default:
		return command.NewErrorArgument(fmt.Errorf("unsupported storage provider: %s", l.Provider))
	}
}
