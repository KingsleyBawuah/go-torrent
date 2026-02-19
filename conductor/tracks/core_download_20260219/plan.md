# Plan: Complete the Core Torrent Download Loop

## Phase 1: Establish Peer Connections
- [ ] Implement `Peer` struct method to initiate TCP connection.
- [ ] Implement `Handshake` method: send handshake and validate response.
- [ ] Test: Mock a peer server and verify handshake success/failure.

## Phase 2: Message Protocol
- [ ] Define `Message` struct for BitTorrent messages (ID, Payload).
- [ ] Implement `ReadMessage` and `WriteMessage` for connection handling.
- [ ] specific message parsing: `Choke`, `Unchoke`, `Interested`, `NotInterested`, `Have`, `Bitfield`, `Request`, `Piece`.
- [ ] Test: Unit tests for serializing/deserializing each message type.

## Phase 3: Bitfield & State Management
- [ ] Create `Bitfield` type to track peer pieces.
- [ ] Implement `HasPiece` and `SetPiece` methods.
- [ ] Update `Client` to manage `Bitfield` for the download progress.
- [ ] Test: Verify bitfield logic with mocked data.

## Phase 4: Download Logic (The Worker)
- [ ] Implement `DownloadWorker`: connects to peer, performs handshake, receives bitfield.
- [ ] Implement `sendInterested` and `recvUnchoke` flow.
- [ ] Implement `downloadPiece`: request blocks, buffer them, validate SHA-1.
- [ ] Test: Simulate full piece download with mocked peer responses.

## Phase 5: Job Queue & Orchestration
- [ ] Create `WorkQueue` to manage pieces (Pending, InProgress, Complete).
- [ ] Implement `Client.Download()`: spawn workers, distribute jobs, handle errors/retries.
- [ ] Integration: Connect `WorkQueue` to `DownloadWorker`.

## Phase 6: File Assembly & Output
- [ ] Implement writing downloaded pieces to disk at correct offsets.
- [ ] Finalize file checksum verification.
- [ ] Test: Verify file integrity after mock download.

## Phase 7: CLI Integration
- [ ] Update `cmd/get.go` to initialize `Client` and start download.
- [ ] Add progress bar/logging for user feedback.
