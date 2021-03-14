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

func NewPeerList(buf []byte) []Peer {
	var peerList []Peer
	chunkedList := helper.ChunkBy(buf, 6)
	for _, val := range chunkedList {
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
