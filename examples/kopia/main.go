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

package main

import (
	"fmt"

	"github.com/kanisterio/safecli/examples/kopia/args"
	"github.com/kanisterio/safecli/examples/kopia/repository"
)

func main() {
	args := repository.CreateArgs{
		Common: args.Common{
			ConfigFilePath: "/path/to/config",
			LogDirectory:   "/path/to/log",
			LogLevel:       "error",
			RepoPassword:   "password",
		},
		Hostname: "localhost",
		Username: "user",
	}
	cmd, err := repository.Create(args)

	fmt.Printf("exec=%#v\n", cmd.Build())
	fmt.Printf("log=%#v\n", cmd) // make sure that password is redacted
	fmt.Printf("log=%v\n", cmd)  // make sure that password is redacted
	fmt.Printf("err=%#v\n", err)
}

// $ go run .
// exec=[]string{"kopia", "--config-file=/path/to/config", "--log-dir=/path/to/log", "--log-level=error", "--password=password", "repository", "create", "--override-hostname=localhost", "--override-username=user"}
// log=&safecli.Builder{Args:[]safecli.Argument{safecli.Argument{Key:"", Value:(*safecli.PlainValue)(0xc000096020)}, safecli.Argument{Key:"--config-file", Value:(*safecli.PlainValue)(0xc000096030)}, safecli.Argument{Key:"--log-dir", Value:(*safecli.PlainValue)(0xc000096040)}, safecli.Argument{Key:"--log-level", Value:(*safecli.PlainValue)(0xc000096050)}, safecli.Argument{Key:"--password", Value:<****>}, safecli.Argument{Key:"", Value:(*safecli.PlainValue)(0xc000096080)}, safecli.Argument{Key:"", Value:(*safecli.PlainValue)(0xc0000960a0)}, safecli.Argument{Key:"--override-hostname", Value:(*safecli.PlainValue)(0xc0000960b0)}, safecli.Argument{Key:"--override-username", Value:(*safecli.PlainValue)(0xc0000960c0)}}, Formatter:(safecli.ArgumentFormatter)(0x47cee0)}
// log=kopia --config-file=/path/to/config --log-dir=/path/to/log --log-level=error --password=<****> repository create --override-hostname=localhost --override-username=user
// err=<nil>
