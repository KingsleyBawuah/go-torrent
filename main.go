package main

import (
	"log"
	"os"
)

func main() {

	log.Println("Welcome to Go-Torrent!")
	torrentfilePath := os.Args[1]
	log.Printf("Is your filepath, %s ? \n", torrentfilePath)

	//Find and Open the file.
	torrentFile, err := os.Open(torrentfilePath)
	if err != nil {
		if err == err.(*os.PathError) {
			log.Panic("Error finding your file. Are you sure this is the correct path?", torrentfilePath)
		}
		log.Panic("Error opening file: ", err)
	}
	defer torrentFile.Close()

	//Read the file
	buffer := make([]byte, 1337)
	_, err = torrentFile.Read(buffer)
	if err != nil {
		log.Panic("Error reading file contents", err)
	}
	log.Print("Here are the contents of your file: \n ", string(buffer))

}
