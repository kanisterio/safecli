# SafeCLI: Secure Command Line Interface Package

SafeCLI is a specialized Go package aimed at facilitating the secure construction, redaction, and logging of command-line arguments. It's designed particularly to handle sensitive values, ensuring they are managed securely throughout the process of command-line argument preparation and execution. This package is part of the Kanister project.

## Features

- **Secure Argument Handling**: Safely handles sensitive information in command-line arguments, ensuring that sensitive details are never exposed in logs or error messages.
- **Flexible CLI Construction**: Provides a builder pattern for constructing command-line arguments dynamically, allowing for clean and readable code.
- **Redaction Interface**: Integrates a redaction system that automatically obscures sensitive information when arguments are converted to strings for logging or debugging.
- **Logging Utility**: Includes a logging utility that ensures sensitive information is redacted, while still providing helpful output for debugging and monitoring.

## Installation

You can install SafeCLI by running:

```bash
go get -u github.com/kanisterio/safecli
```

## Usage

### Basic CLI Construction

To construct a CLI command, you can use the `NewBuilder` function which initializes a new command builder. You can append arguments, both loggable and sensitive, to the builder:

```go
import "github.com/kanisterio/safecli"

func main() {
    builder := safecli.NewBuilder("zip").
        AppendLoggableKV("--output", "/path/to/output").
        AppendRedactedKV("--password", "secretPass").
        AppendLoggable("input_file1", "input_file2")

    command := builder.Build()
    // Use `command` as required, e.g., execute it
}
```

In the above example, `--password` is a sensitive argument, and its value will be redacted in logs.

### Secure Logging

To log a CLI command securely, ensuring that sensitive information is redacted:

```go
logger := safecli.NewLogger(builder)
logOutput := logger.Log()
// Use `logOutput` for logging
```

## Contribution

Contributions to the SafeCLI package are welcome. Please follow the contribution guidelines of the Kanister project when submitting patches or features.
