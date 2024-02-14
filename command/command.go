package command

import "github.com/kanisterio/safecli"

// New returns a new safecli.Builder for the storage sub command.
func New(name string, args ...Applier) (*safecli.Builder, error) {
	cmd := safecli.NewBuilder(name)
	if err := Apply(cmd, args...); err != nil {
		return nil, err
	}
	return cmd, nil
}
