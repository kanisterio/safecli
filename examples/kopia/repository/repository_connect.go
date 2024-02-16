package repository

import (
	"time"

	"github.com/kanisterio/safecli"
	"github.com/kanisterio/safecli/examples/kopia/args"
	"github.com/kanisterio/safecli/examples/kopia/internal"
	"github.com/kanisterio/safecli/examples/kopia/internal/opts"
)

// ConnectArgs defines the arguments for the `kopia repository connect` command.
type ConnectArgs struct {
	args.Common           // common arguments
	Location    Location  // filesystem, s3, etc
	Hostname    string    // the hostname of the repository
	Username    string    // the username of the repository
	ReadOnly    bool      // connect to a repository in read-only mode
	PointInTime time.Time // connect to a repository as it was at a specific point in time
}

// Connect creates a new safecli.Builder for the `kopia repository connect` command.
func Connect(args ConnectArgs) (*safecli.Builder, error) {
	return internal.NewKopiaCommand(opts.Common(args.Common),
		cmdRepository, subcmdConnect,
		optHostname(args.Hostname),
		optUsername(args.Username),
		optReadOnly(args.ReadOnly),
		optPointInTime(args.PointInTime),
		optStorage(args.Location),
	)
}
