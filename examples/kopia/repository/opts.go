package repository

import "github.com/kanisterio/safecli/command"

// cmdXXX represents different `kopia repository“ commands.
var (
	cmdRepository = command.NewArgument("repository")
	cmdCreate     = command.NewArgument("create")
)

// optHostname returns a new optHostname flag with a given optHostname.
func optHostname(hostname string) command.Applier {
	return command.NewOptionWithArgument("--override-hostname", hostname)
}

// optUsername returns a new optUsername flag with a given optUsername.
func optUsername(username string) command.Applier {
	return command.NewOptionWithArgument("--override-username", username)
}
