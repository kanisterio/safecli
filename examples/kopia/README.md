# Kopia Command Line Interface(CLI) Builder

This example demonstrates how to use [safecli](https://github.com/kanisterio/safecli) for programmatically building a CLI for [Kopia](https://github.com/kopia/kopia), focusing on `kopia repository create` and `kopia repository connect` commands.

## Building the CLI

### Common Arguments

Kopia repository commands share a set of common arguments, such as `--config-file`, `--log-dir`, `--log-level`, and `--password`. These common arguments are defined in the [args](args/common_args.go) package and are utilized in both `create` and `connect` commands.

### Repository Commands

The commands related to repositories are contained within the [examples/kopia/repository](repository/) package.


#### Define Repository Creation Arguments

First, define the arguments for the `kopia repository create` command as [repository.CreateArgs](repository/repository_create.go) structure, which embeds the common arguments and adds the `Location`, `Hostname`, and `Username` fields:

```go
// CreateArgs represents the arguments for the `kopia repository create` command.
type CreateArgs struct {
    args.Common          // Embeds common arguments
    Location    Location // Filesystem, S3, etc.
    Hostname    string   // The hostname of the repository
    Username    string   // The username of the repository
}
```

#### Define the Repository Creation Function

Next, define the [repository.Create](repository/repository_create.go) function to create a new [safecli.Builder](https://github.com/kanisterio/safecli/blob/main/safecli.go) for the command using arguments from `CreateArgs` structure:

```go
// Create creates a new safecli.Builder for the `kopia repository create` command.
func Create(args CreateArgs) (*safecli.Builder, error) {
    return internal.NewKopiaCommand(
        opts.Common(args.Common),
        cmdRepository, subcmdCreate,
        optHostname(args.Hostname),
        optUsername(args.Username),
        optStorage(args.Location),
    )
}
```

This function calls `internal.NewKopiaCommand` from the [examples/kopia/internal](internal/kopia.go) package to create a `safecli.Builder`, converting `CreateArgs` to CLI options through `opts.Common`, `optHostname`, `optUsername`, and `optStorage`.

Common options are defined in [examples/kopia/internal/opts/common_opts.go](internal/opts/common_opts.go) and repository options are defined in [examples/kopia/repository/opts.go](repository/opts.go) files.

#### Example Usage

To build the `kopia repository create ...` command from your Go code, you must use the `Create` function:

```go
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
    fmt.Printf("exec=%#v\n", cmd.Build())
}
```

This code will print the command that you can run to create a new Kopia repository.

```bash
$ go run main.go
exec=[]string{"kopia", "--config-file=/path/to/config", "--log-dir=/path/to/log", "--log-level=error", "--password=123456", "repository", "create", "--override-hostname=localhost", "--override-username=user", "filesystem", "--path=/tmp/my-repository"}
```

#### Repository Connect command 

The `repository connect` command is implemented in a similar way. You can find the complete example in the [examples/kopia/repository_connect.go](repository/repository_connect.go).

Usage example can be found in [examples/kopia/main.go](main.go).


## Bottom Line

This example demonstrates how to use `safecli` to programmatically build a CLI for Kopia, focusing on the `kopia repository create` and `kopia repository connect` commands. The same approach can be applied to construct other Kopia commands or any other CLI tool.
