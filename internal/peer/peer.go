package peer

import (
	"bytes"
	"encoding/binary"
	"github.com/KingsleyBawuah/go-torrent/internal/helper"
	"log"
	"net"
)

type Peer struct {
	PeerId string `bencode:"peer id"`
	Ip     string `bencode:"ip"`
	Port   int64  `bencode:"port"`
}

// NewPeerList creates a slice of type Peer from a slice of bytes that correspond to the binary peer list model.
func NewPeerList(buf []byte) []Peer {
	var peerList []Peer
	chunkedList := helper.ChunkBy(buf, 6)
	for _, val := range chunkedList {
		//First 4 bytes are the IP addr, final two are the port number.
		var port uint16
		b := bytes.NewReader(val[4:])
		err := binary.Read(b, binary.BigEndian, &port)
		if err != nil {
			log.Println("binary.Read failed:", err)
		}

		peerList = append(peerList, Peer{
			Ip:   net.IP(val[0:4]).String(),
			Port: int64(port),
		})
	}
	return peerList
}
