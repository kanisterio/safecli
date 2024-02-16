package repository

// Provider represents the storage provider for the repository.
type Provider string

const (
	ProviderFilesystem Provider = "filesystem"
	ProviderS3         Provider = "s3"
)

// Location represents the location of the repository.
type Location struct {
	Provider Provider
	MetaData map[string][]byte
}
