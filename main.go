package main

import (
	"bytes"
	"github.com/zeebo/bencode"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.Println("Welcome to go-torrent!")
	torrentFilepath := os.Args[1]

	//Find and Open the file.
	torrentFile, err := os.Open(torrentFilepath)
	if err != nil {
		if err == err.(*os.PathError) {
			log.Panic("Error finding your file. Are you sure this is the correct path?", torrentFilepath)
		}
		log.Panic("Error opening file: ", err)
	}
	defer torrentFile.Close()

	//Read the file
	torrentFileContent, err := ioutil.ReadFile(torrentFilepath)
	if err != nil {
		log.Panic("Error reading file: \n", err)
	}

	//Information on torrent file structure: https://wiki.theory.org/BitTorrentSpecification
	type torrentInfoBencode struct {
		Pieces      string `bencode:"pieces"`       //string consisting of the concatenation of all 20-byte SHA1 hash values, one per piece.
		PieceLength int64  `bencode:"piece length"` //Bytes per piece
	}

	type torrent struct {
		Announce     string             `bencode:"announce"`      //Url of the tracker
		Comment      string             `bencode:"comment"`       //Comment left on the torrent. Optional
		CreationDate int64              `bencode:"creation date"` //Date torrent was created
		CreatedBy    string             `bencode:"created by"`    //Author. Optional
		Encoding     string             `bencode:"encoding"`      //the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile. Optional
		Info         torrentInfoBencode `bencode:"info"`          //TODO: Handle multiple file case.
	}

	inputTorrentFile := &torrent{}

	//Decode the file.
	r := bytes.NewReader(torrentFileContent)
	decoder := bencode.NewDecoder(r)

	if err := decoder.Decode(inputTorrentFile); err != nil {
		log.Panic("Error decoding torrent file: ", err)
	}

	log.Printf("Decoded torrent file %+v\\n", inputTorrentFile)
}
