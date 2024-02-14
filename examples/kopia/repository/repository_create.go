package repository

import (
	"github.com/kanisterio/safecli"
	"github.com/kanisterio/safecli/examples/kopia/args"
	"github.com/kanisterio/safecli/examples/kopia/internal"
	"github.com/kanisterio/safecli/examples/kopia/internal/opts"
)

type CreateArgs struct {
	args.Common

	Hostname string
	Username string
}

func Create(args CreateArgs) (*safecli.Builder, error) {
	return internal.NewKopiaCommand(opts.Common(args.Common),
		cmdRepository, cmdCreate,
		optHostname(args.Hostname),
		optUsername(args.Username),
	)
}
