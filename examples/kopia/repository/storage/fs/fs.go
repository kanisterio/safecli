package fs

import (
	"github.com/kanisterio/safecli/command"
)

// metadata is the metadata for the filesystem storage.
type metadata map[string][]byte

func (f metadata) RepoPath() string {
	return string(f["repoPath"])
}

// New creates a new subcommand for the filesystem storage.
func New(data map[string][]byte) command.Applier {
	m := metadata(data)
	return command.NewArguments(subcmdFilesystem,
		optRepoPath(m.RepoPath()),
	)
}
