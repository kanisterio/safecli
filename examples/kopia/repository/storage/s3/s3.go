package s3

import (
	"strconv"

	"github.com/kanisterio/safecli/command"
)

// metadata is the metadata for the S3 storage.
type metadata map[string][]byte

func (f metadata) get(key string) string {
	return string(f[key])
}

func (f metadata) Region() string {
	return f.get("region")
}

func (f metadata) BucketName() string {
	return f.get("bucket")
}

func (f metadata) Endpoint() string {
	return f.get("endpoint")
}

func (f metadata) Prefix() string {
	return f.get("prefix")
}

func (f metadata) IsInsecureEndpoint() bool {
	return f.get("endpoint") == "http"
}

func (f metadata) HasSkipSSLVerify() bool {
	v, _ := strconv.ParseBool(f.get("skipSSLVerify"))
	return v
}

// New creates a new subcommand for the S3 storage.
func New(data map[string][]byte) command.Applier {
	m := metadata(data)
	return command.NewArguments(subcmdS3,
		optRegion(m.Region()),
		optBucket(m.BucketName()),
		optEndpoint(m.Endpoint()),
		optPrefix(m.Prefix()),
		optDisableTLS(m.IsInsecureEndpoint()),
		optDisableTLSVerify(m.HasSkipSSLVerify()),
	)
}
