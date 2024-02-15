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
			RepoPassword:   "123456",
		},
		Location: repository.Location{
			Provider: "filesystem",
			MetaData: map[string][]byte{
				"repoPath": []byte("/tmp/my-repository"),
			},
		},
		Hostname: "localhost",
		Username: "user",
	}
	cmd, err := repository.Create(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("exec=", cmd)
	// exec= kopia --config-file=/path/to/config --log-dir=/path/to/log --log-level=error --password=<****> repository create --override-hostname=localhost --override-username=user filesystem --path=/tmp/my-repository
	fmt.Printf("exec=%#v", cmd.Build())
	// exec= kopia --config-file=/path/to/config --log-dir=/path/to/log --log-level=error --password=<****> repository create --override-hostname=localhost --override-username=user filesystem --path=/tmp/my-repository
}

// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/kanisterio/safecli/examples/kopia/args"
// 	"github.com/kanisterio/safecli/examples/kopia/repository"
// )

// func fsCreateArgs() repository.CreateArgs {
// 	return repository.CreateArgs{
// 		Common: args.Common{
// 			ConfigFilePath: "/path/to/config",
// 			LogDirectory:   "/path/to/log",
// 			LogLevel:       "error",
// 			RepoPassword:   "password",
// 		},
// 		Location: repository.Location{
// 			Provider: repository.ProviderFilesystem,
// 			MetaData: map[string][]byte{
// 				"repoPath": []byte("/tmp/my-repository"),
// 			},
// 		},
// 		Hostname: "localhost",
// 		Username: "user",
// 	}
// }

// func fsConnectArgs() repository.ConnectArgs {
// 	return repository.ConnectArgs{
// 		Common: args.Common{
// 			ConfigFilePath: "/path/to/config",
// 			LogDirectory:   "/path/to/log",
// 			LogLevel:       "error",
// 			RepoPassword:   "password",
// 		},
// 		Location: repository.Location{
// 			Provider: repository.ProviderFilesystem,
// 			MetaData: map[string][]byte{
// 				"repoPath": []byte("/tmp/my-repository"),
// 			},
// 		},
// 		Hostname:    "localhost",
// 		Username:    "user",
// 		ReadOnly:    true,
// 		PointInTime: time.Date(2024, 2, 15, 14, 30, 0, 0, time.FixedZone("PST", -8*60*60)),
// 	}
// }

// func s3CreateArgs() repository.CreateArgs {
// 	return repository.CreateArgs{
// 		Common: args.Common{
// 			ConfigFilePath: "/path/to/config",
// 			LogDirectory:   "/path/to/log",
// 			LogLevel:       "error",
// 			RepoPassword:   "password",
// 		},
// 		Location: repository.Location{
// 			Provider: repository.ProviderS3,
// 			MetaData: map[string][]byte{
// 				"region":        []byte("us-west-1"),
// 				"bucket":        []byte("my-bucket"),
// 				"prefix":        []byte("my-repository"),
// 				"endpoint":      []byte("http://localhost:9000"),
// 				"skipSSLVerify": []byte("true"),
// 			},
// 		},
// 		Hostname: "localhost",
// 		Username: "user",
// 	}
// }

// func RepoCreate(args repository.CreateArgs) {
// 	cmd, err := repository.Create(args)
// 	fmt.Println("exec=", cmd)
// 	fmt.Println("err=", err)
// }

// func RepoConnect(args repository.ConnectArgs) {
// 	cmd, err := repository.Connect(args)
// 	fmt.Println("exec=", cmd)
// 	fmt.Println("err=", err)
// }

// func main() {
// 	RepoCreate(fsCreateArgs())
// 	RepoCreate(s3CreateArgs())
// 	RepoConnect(fsConnectArgs())
// }

// // $ go run main.go
// // exec= kopia --config-file=/path/to/config --log-dir=/path/to/log --log-level=error --password=<****> repository create --override-hostname=localhost --override-username=user filesystem --path=/tmp/my-repository
// // err= <nil>
// // exec= kopia --config-file=/path/to/config --log-dir=/path/to/log --log-level=error --password=<****> repository create --override-hostname=localhost --override-username=user s3 --region=us-west-1 --bucket=my-bucket --endpoint=http://localhost:9000 --prefix=my-repository --disable-tls-verify
// // err= <nil>
// // exec= kopia --config-file=/path/to/config --log-dir=/path/to/log --log-level=error --password=<****> repository connect --override-hostname=localhost --override-username=user --read-only --point-in-time=2024-02-15T14:30:00-08:00 filesystem --path=/tmp/my-repository
// // err= <nil>
