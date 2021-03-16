package peer

import (
	"reflect"
	"testing"
)

func TestNewPeerList(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want []Peer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitSwarm(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitSwarm() = %v, want %v", got, tt.want)
			}
		})
	}
}
