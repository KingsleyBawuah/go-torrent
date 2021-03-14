package tracker

import (
	"github.com/KingsleyBawuah/go-torrent/internal/metainfo"
	"github.com/zeebo/bencode"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Tracker struct {
	Url      url.URL //The url location of the tracker.
	InfoHash string  //SHA-1 Hash of the bencoded Info field located in a metainfo file.
}

// Response represents the bencode dictionary response from a tracker.
type Response struct {
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

// New creates a new Tracker.
func New(metainfo metainfo.HashManager, link string) Tracker {
	u, err := url.Parse(link)
	if err != nil {
		log.Panic("Error parsing tracker url.")
	}
	return Tracker{
		Url:      *u,
		InfoHash: metainfo.InfoHash(),
	}
}

//Construct and make the initial request to the tracker.
func (t Tracker) Req() Response {
	//Construct a request to the tracker
	trackerReq := url.Values{}
	trackerReq.Add("info_hash", t.InfoHash)
	trackerReq.Add("peer_id", "GT-19362856256926571") //TODO: Figure out how I'm going to generate these.
	trackerReq.Add("port", "6881")
	trackerReq.Add("uploaded", "0")
	trackerReq.Add("downloaded", "0")
	trackerReq.Add("left", "1000")
	trackerReq.Add("compact", "0")
	trackerReq.Add("no_peer_id", "0")
	trackerReq.Add("event", "started")

	trackerUrl, err := url.Parse(t.Url.String())
	if err != nil {
		log.Panic("Error parsing tracker url: ", err)
	}
	trackerUrl.RawQuery = trackerReq.Encode()

	log.Print("Requesting information from the tracker at: ", trackerUrl.String())

	resp, err := http.Get(trackerUrl.String())
	if err != nil {
		log.Panicf("Error getting information from the torrent's announce url: %s. Error message: %s", trackerUrl.String(), err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//Decode the tracker response
	bodyBencode := Response{}
	if err := bencode.DecodeBytes(body, &bodyBencode); err != nil {
		log.Panic("Error decoding tracker response: ", err)
	}

	return bodyBencode
}
