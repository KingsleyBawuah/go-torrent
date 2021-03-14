/*
Copyright © 2021 Kingsley Bawuah <Kingsley404@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"crypto/sha1"
	"github.com/KingsleyBawuah/go-torrent/internal/metainfo"
	"github.com/KingsleyBawuah/go-torrent/internal/peer"
	"github.com/KingsleyBawuah/go-torrent/internal/tracker"
	"github.com/spf13/cobra"
	"github.com/zeebo/bencode"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Begin downloading from a .torrent file.",
	Long: `Get begins the process of torrenting a file. 
	It must be provided the absolute path to an existing .torrent file.
	
	For example:

	go-torrent get /Users/root/Downloads/debian-10.8.0-amd64-netinst.iso.torrent
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Separate this logic into it's own file.
		log.Println("Welcome to go-torrent!")
		torrentFilepath := args[0]

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
			log.Panic("Error reading file: ", err)
		}
		inputTorrentFile := &metainfo.Torrent{}

		//Determine if we are downloading a single file or multiple.
		singleFileMode := inputTorrentFile.IsSingleFile()
		if singleFileMode {
			log.Println("Single file mode.....")
		} else {
			log.Println("Multi file mode.....")
		}

		//Decode the file.
		r := bytes.NewReader(torrentFileContent)
		decoder := bencode.NewDecoder(r)

		if err := decoder.Decode(inputTorrentFile); err != nil {
			log.Panic("Error decoding torrent file: ", err)
		}

		if inputTorrentFile.Announce == "" {
			log.Panic("The torrent file you supplied does not have an announce field. This client only supports the original BitTorrent Spec. Please try another torrent.")
		}

		//SHA-1 Encode the Info key's value.
		infoBytes, err := bencode.EncodeBytes(inputTorrentFile.Info)
		if err != nil {
			log.Panic("Error encoding bytes for info field.")
		}
		h := sha1.New()
		h.Write(infoBytes)
		infoHash := h.Sum(nil)

		//Construct a request to the tracker
		trackerReq := url.Values{}
		trackerReq.Add("info_hash", string(infoHash))
		trackerReq.Add("peer_id", "GT-19362856256926571") //TODO: Figure out how I'm going to generate these.
		trackerReq.Add("port", "6881")
		trackerReq.Add("uploaded", "0")
		trackerReq.Add("downloaded", "0")
		trackerReq.Add("left", "1000")
		trackerReq.Add("compact", "0")
		trackerReq.Add("no_peer_id", "0")
		trackerReq.Add("event", "started")

		trackerUrl, err := url.Parse(inputTorrentFile.Announce)
		if err != nil {
			log.Panic("Error parsing tracker url: ", err)
		}
		trackerUrl.RawQuery = trackerReq.Encode()

		log.Print("Requesting information from the tracker at: ", trackerUrl.String())

		resp, err := http.Get(trackerUrl.String())
		if err != nil {
			log.Panicf("Error getting information from the torrent's announce url: %s. Error message: %s", inputTorrentFile.Announce, err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		//Decode the tracker response
		bodyBencode := tracker.Response{}
		if err := bencode.DecodeBytes(body, &bodyBencode); err != nil {
			log.Panic("Error decoding tracker response: ", err)
		}

		log.Print("Peer list from the tracker response: ", peer.NewPeerList(bodyBencode.Peers))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
