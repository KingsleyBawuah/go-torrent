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
		{
			name: "Testing even numbered array, even chunk size",
			args: args{
				items:     []byte{128, 107, 163, 83},
				chunkSize: 2,
			},
			wantChunks: [][]byte{{128, 107}, {163, 83}},
		},
		{
			name: "Testing even numbered array, odd chunk size",
			args: args{
				items:     []byte{128, 107, 163, 83},
				chunkSize: 3,
			},
			wantChunks: [][]byte{{128, 107, 163}, {83}},
		},

		{
			name: "Testing odd numbered array, odd chunk size",
			args: args{
				items:     []byte{128, 107, 163, 83, 188},
				chunkSize: 3,
			},
			wantChunks: [][]byte{{128, 107, 163}, {83, 188}},
		},

		{
			name: "Testing odd numbered array, even chunk size",
			args: args{
				items:     []byte{128, 107, 163, 83, 188},
				chunkSize: 2,
			},
			wantChunks: [][]byte{{128, 107}, {163, 83}, {188}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChunks := ChunkBy(tt.args.items, tt.args.chunkSize); !reflect.DeepEqual(gotChunks, tt.wantChunks) {
				t.Errorf("ChunkBy() = %v, want %v", gotChunks, tt.wantChunks)
			}
		})
	}
}
