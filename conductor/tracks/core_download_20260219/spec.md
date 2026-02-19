# Specification: Core Torrent Download Loop

## Goal
Implement the core logic to download a single-file torrent from start to finish, adhering to the BitTorrent Protocol Specification (BEP 0003).

## Scope
- **Connect to Peers:** Establish TCP connections with peers discovered from the tracker.
- **Handshake:** Perform the standard BitTorrent handshake.
- **Message Protocol:** Implement reading and writing of standard messages:
    - `choke`, `unchoke`, `interested`, `not interested`, `have`, `bitfield`, `request`, `piece`, `cancel`.
- **Piece Management:**
    - Track which pieces we have and which we need.
    - Validate downloaded pieces against the SHA-1 hash in the `.torrent` file.
- **Pipelining:** Request blocks in chunks (16KB) to maximize throughput.
- **File Writing:** Write verified pieces to disk.

## Technical Details

### 1. Peer Protocol
- **Handshake:**
    - `pstrlen`: 19
    - `pstr`: "BitTorrent protocol"
    - `reserved`: 8 bytes (all zeros for now)
    - `info_hash`: 20 bytes (from metainfo)
    - `peer_id`: 20 bytes (our client ID)
- **Message Format:** `<length prefix><message ID><payload>`

### 2. Job Queue (Piece Manager)
- A thread-safe queue of pieces to download.
- Workers (peer connections) pick pieces from the queue.
- Failed pieces (due to disconnect or hash mismatch) are returned to the queue.

### 3. Client Structure
- `Client`: Orchestrates the download.
    - Maintains connection to peers.
    - Manages the `WorkQueue`.
    - Handles writing to the output file.

## Out of Scope (for this track)
- Multi-file torrents (partially handled in parsing, but download logic will focus on single file first).
- Seeding (we will focus on leeching/downloading first).
- DHT / Magnet links.
