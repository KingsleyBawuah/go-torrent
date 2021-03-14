package tracker

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
