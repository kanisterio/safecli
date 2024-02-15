package s3

import (
	"fmt"

	"github.com/kanisterio/safecli/command"
)

var (
	subcmdS3 = command.NewArgument("s3")
)

// optRegion creates a new region option with a given region.
// if the region is empty, it will do nothing.
func optRegion(region string) command.Applier {
	return command.NewOptionWithArgument("--region", region)
}

// optBucket creates a new bucket option with a given name.
// It returns an error if the name is empty.
func optBucket(name string) command.Applier {
	if name == "" {
		return command.NewErrorArgument(fmt.Errorf("bucket name cannot be empty"))
	}
	return command.NewOptionWithArgument("--bucket", name)
}

// optEndpoint creates a new endpoint option with a given endpoint.
// if the endpoint is empty, it will do nothing.
func optEndpoint(endpoint string) command.Applier {
	return command.NewOptionWithArgument("--endpoint", endpoint)
}

// optPrefix creates a new prefix option with a given prefix.
// if the prefix is empty, it will do nothing.
func optPrefix(prefix string) command.Applier {
	return command.NewOptionWithArgument("--prefix", prefix)
}

// optDisableTLS creates a new disable-tls option with a given value.
// if the disable is false, it will do nothing.
func optDisableTLS(disable bool) command.Applier {
	return command.NewOption("--disable-tls", disable)
}

// optDisableTLSVerify creates a new disable-tls-verify option with a given value.
// if the disable is false, it will do nothing.
func optDisableTLSVerify(disable bool) command.Applier {
	return command.NewOption("--disable-tls-verify", disable)
}
