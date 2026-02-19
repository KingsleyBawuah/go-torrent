# go-torrent

`go-torrent` is a CLI BitTorrent client written in Go. It aims to implement the original [BitTorrent Protocol Specification](https://tixati.com/specs/bittorrent).

## Project Overview

- **Language:** Go (1.15+)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Configuration:** [Viper](https://github.com/spf13/viper)
- **Bencode Parsing:** [zeebo/bencode](https://github.com/zeebo/bencode)

### Architecture

The project follows a standard Go project layout:

- `main.go`: Entry point of the application.
- `cmd/`: Contains CLI command definitions using Cobra.
  - `root.go`: Base command and configuration initialization.
  - `get.go`: The `get` command for initiating downloads from `.torrent` files.
- `internal/`: Core business logic, separated into packages:
  - `metainfo/`: Handles parsing of `.torrent` files and generation of the `InfoHash`.
  - `tracker/`: Manages communication with BitTorrent trackers.
  - `peer/`: Handles peer discovery and swarm initialization.
  - `helper/`: Utility functions (e.g., slice chunking).

## Building and Running

### Prerequisites

- Go 1.15 or higher.

### Commands

- **Build:**
  ```bash
  go build -o go-torrent main.go
  ```

- **Run:**
  ```bash
  ./go-torrent get <path-to-torrent-file>
  ```

- **Test:**
  ```bash
  go test ./...
  ```

## Development Conventions

- **Internal Logic:** All core domain logic should reside within the `internal/` directory to prevent external usage as a library if not intended.
- **Protocol Adherence:** Focus on BEP 0003 specification.
- **Error Handling:** Currently uses `log.Panic` for critical errors; transitioning to more idiomatic Go error handling is a likely future improvement.
- **Testing:** Unit tests are located alongside the source code with `_test.go` suffix (e.g., `internal/metainfo/torrent_test.go`).
