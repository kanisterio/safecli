package command

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
//

//
// The command package is used to define CLI (Command Line Interface) commands along with their arguments.
//
// Command line arguments are the whitespace-separated tokens given in the shell command used to invoke the program.
//
// A token prefixed with a hyphen delimiter (`-`) is known as an *option*. For example, `-o` or `--option`.
//
// An option may or may not have an associated argument. For example, `--option=value`.
//
// A token without a hyphen delimiter (`-`) is considered an *argument*. For example, `arg1` or `arg2`.
//
// The command package provides a set of interfaces and types for defining and applying arguments to commands.
//
// Check safecli/examples/kopia package for usage of the command package.
//
