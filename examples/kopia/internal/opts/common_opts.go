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

package opts

import (
	"github.com/kanisterio/safecli/command"
	"github.com/kanisterio/safecli/examples/kopia/args"
)

// defaults
var (
	defaultLogLevel = "error"
)

// ConfigFilePath returns a new config file path flag with a given path.
func ConfigFilePath(path string) command.Applier {
	return command.NewOptionWithArgument("--config-file", path)
}

// LogDirectory returns a new log directory flag with a given directory.
func LogDirectory(dir string) command.Applier {
	return command.NewOptionWithArgument("--log-dir", dir)
}

// LogLevel returns a new log level flag with a given level.
func LogLevel(level string) command.Applier {
	if level == "" {
		level = defaultLogLevel
	}
	return command.NewOptionWithArgument("--log-level", level)
}

// RepoPassword returns a new repository password flag with a given password.
func RepoPassword(password string) command.Applier {
	return command.NewOptionWithRedactedArgument("--password", password)
}

// Common maps the common arguments to the CLI command arguments.
func Common(args args.Common) command.Applier {
	return command.NewArguments(
		ConfigFilePath(args.ConfigFilePath),
		LogDirectory(args.LogDirectory),
		LogLevel(args.LogLevel),
		RepoPassword(args.RepoPassword),
	)
}
