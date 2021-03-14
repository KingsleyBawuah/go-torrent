package peer

import (
	"bytes"
	"encoding/binary"
	"github.com/KingsleyBawuah/go-torrent/internal/helper"
	"log"
	"net"
)

type TrackerResponse struct {
	FailureReason  string `bencode:"failure reason"`  //Human readable error as to why req failed.
	WarningMessage string `bencode:"warning message"` //Warning, request should process as usual
	Interval       int64  `bencode:"interval"`        //Seconds between the requests to tracker.
	MinInterval    string `bencode:"min interval"`    //If present clients must not re-announce more frequently than this.
	TrackerId      string `bencode:"tracker id"`      //String that should be sent back on next announcement
	Seeders        int64  `bencode:"complete"`        //Number of peers with entire file.
	Leechers       int64  `bencode:"incomplete"`      //Number of peers not seeding.
	//TODO: Handle cases in which the peers field is in regular bencode form vs byte array.
	Peers []byte `bencode:"peers"` //Contains information about peers.
}

type Peer struct {
	PeerId string `bencode:"peer id"`
	Ip     string `bencode:"ip"`
	Port   int64  `bencode:"port"`
}

func (t *TrackerResponse) PeerList() []Peer {
	var peerList []Peer
	chunkedList := helper.ChunkBy(t.Peers, 6)
	for _, val := range chunkedList {
		var port uint16

		buf := bytes.NewReader(val[4:])
		err := binary.Read(buf, binary.BigEndian, &port)
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
