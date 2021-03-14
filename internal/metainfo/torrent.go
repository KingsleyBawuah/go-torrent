package metainfo

//Represents a .torrent metainfo file.
//More information on torrent file structure: https://wiki.theory.org/BitTorrentSpecification
type Torrent struct {
	Announce     string `bencode:"announce"`      //Url of the tracker
	Comment      string `bencode:"comment"`       //Comment left on the torrent. Optional
	CreationDate int64  `bencode:"creation date"` //Date torrent was created
	CreatedBy    string `bencode:"created by"`    //Author. Optional
	Encoding     string `bencode:"encoding"`      //the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile. Optional
	Info         Info   `bencode:"info"`          //Dictionary containing information about the file we want to download.
}

//Contains information about the file we want to download.
type Info struct {
	//In single file mode, Name is the name of the file to be downloaded.
	//In multiple file mode, Name is the name of the directory to store files.
	Name        string  `bencode:"name"`
	Pieces      string  `bencode:"pieces"`            //string consisting of the concatenation of all 20-byte SHA1 hash values, one per piece.
	PieceLength int64   `bencode:"piece length"`      //Bytes per piece
	Private     bool    `bencode:"private,omitempty"` //Optional field may be read as "no external peer source". TODO: Support this?
	Length      int64   `bencode:"length,omitempty"`  //Length of the file to be downloaded, not used in multiple file mode.
	Files       []Files `bencode:"files,omitempty"`   //List of files to be downloaded, not used in single file mode.
}

//Multiple File mode.
type Files struct {
	Length int64    `bencode:"length"`
	Md5sum string   `bencode:"md5sum"`
	Path   []string `bencode:"path"`
}

func (t Torrent) IsSingleFile() bool {
	if t.Info.Files == nil {
		return true
	}
	return false
}
