package internal

import (
	"github.com/kanisterio/safecli"
	"github.com/kanisterio/safecli/command"
)

const (
	kopiaBinName = "kopia"
)

// NewKopiaCommand creates a new safecli.Builder for the kopia command.
func NewKopiaCommand(args ...command.Applier) (*safecli.Builder, error) {
	return command.New(kopiaBinName, args...)
}
