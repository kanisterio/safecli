package fs

import (
	"fmt"

	"github.com/kanisterio/safecli/command"
)

var (
	subcmdFilesystem = command.NewArgument("filesystem")
)

// optRepoPath creates a new path option with a given repoPath.
func optRepoPath(repoPath string) command.Applier {
	if repoPath == "" {
		return command.NewErrorArgument(fmt.Errorf("repoPath cannot be empty"))
	}
	return command.NewOptionWithArgument("--path", repoPath)
}
