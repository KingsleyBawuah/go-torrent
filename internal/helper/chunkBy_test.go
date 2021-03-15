package helper

import (
	"reflect"
	"testing"
)

func TestChunkBy(t *testing.T) {
	type args struct {
		items     []byte
		chunkSize int
	}
	tests := []struct {
		name       string
		args       args
		wantChunks [][]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChunks := ChunkBy(tt.args.items, tt.args.chunkSize); !reflect.DeepEqual(gotChunks, tt.wantChunks) {
				t.Errorf("ChunkBy() = %v, want %v", gotChunks, tt.wantChunks)
			}
		})
	}
}
