# Product Definition: go-torrent

## Initial Concept
A CLI BitTorrent client written in Go. For now this will only implement the original [BitTorrent Protocol Specification](https://tixati.com/specs/bittorrent)

## Vision
`go-torrent` is a modern, lightweight, and educational CLI BitTorrent client. It aims to provide a high-performance downloading experience while serving as a reference implementation for the BitTorrent protocol (v1 and v2) in Go.

## Target Users
- **Resource-Conscious Users:** Individuals seeking a BitTorrent client with a minimal memory footprint and high efficiency.
- **Developers & Students:** Those looking for a readable, well-documented codebase to learn about network protocols and Go-based systems programming.
- **CLI Enthusiasts:** Users who prefer a terminal-based workflow with modern interactive features.

## Core Features
- **Comprehensive Torrent Support:** Support for both single-file and multi-file torrent downloads.
- **Protocol Adherence:** Strict implementation of the BitTorrent Protocol Specification, with a focus on modern standards including BitTorrent v2 (BEP 52).
- **Download Management:** Essential resume capabilities to checkpoint and continue downloads across sessions.
- **Modern CLI UX:** Real-time feedback through progress bars and interactive status updates.

## Technical & Quality Goals
- **Spec Adherence:** Full compliance with BEP 0003 and implementation of modern BEPs like BEP 52 (SHA-256 hashing).
- **Performance:** Optimized for speed and low resource usage.
- **Code Quality:** Maintaining >80% unit test coverage and strict adherence to idiomatic Go standards (`golint`, `go vet`).
- **Readability:** A clean, modular architecture that is easy to navigate and understand.
