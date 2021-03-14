package helper

// Chunkby breaks a slice up into multiple "chunks" aka nested slices with each element containing another slice.
// Used to split up binary list responses.
func ChunkBy(items []byte, chunkSize int) (chunks [][]byte) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
