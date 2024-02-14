package args

// Common represents the common arguments for Kopia commands.
type Common struct {
	ConfigFilePath string // the path to the config file.
	LogDirectory   string // the directory where logs are stored.
	LogLevel       string // the level of logging. Default is "error".
	RepoPassword   string // the password for the repository.
}

// Cache represents the cache arguments for Kopia commands.
type Cache struct {
	// ...
}
