# Technology Stack: go-torrent

## Core Technologies
- **Language:** Go (Targeting version 1.26+)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) - Used for building the CLI interface and command structure.
- **Configuration:** [Viper](https://github.com/spf13/viper) - Handles configuration file loading and management.
- **Bencode Parsing:** [zeebo/bencode](https://github.com/zeebo/bencode) - For parsing and encoding Bencode data used in `.torrent` files and tracker communication.

## Project Structure & Architecture
- **Standard Go Layout:**
    - `cmd/`: CLI entry points and command definitions.
    - `internal/`: Core business logic, protocol implementation, and internal packages (e.g., `metainfo`, `tracker`, `peer`).
- **Package-Based Logic:** Separation of concerns using dedicated internal packages for different parts of the BitTorrent protocol.

## Testing & Quality Assurance
- **Go Testing:** Built-in `testing` package for unit and integration tests.
- **CI/Linting:** Targeting idiomatic Go standards using tools like `golint`, `go vet`, and `go fmt`.
