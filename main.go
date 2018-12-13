package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	id3 "github.com/mikkyang/id3-go"

	"github.com/h2non/filetype"
)

func main() {
	folder := flag.String("folder", "", "[Required] Folder where the audio files are stored")
	album := flag.String("album", "", "Album to set")
	artist := flag.String("artist", "", "Artist to set")
	flag.Parse()

	if flag.NFlag() < 2 || *folder == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	filepath.Walk(*folder, func(path string, info os.FileInfo, err error) error {

		if HasAudioExtension(path) {
			buf, _ := ioutil.ReadFile(path)
			if filetype.IsAudio(buf) {
				fmt.Println(path)
				audioFile, _ := id3.Open(path)
				defer audioFile.Close()

				if *artist != "" {
					audioFile.SetArtist(*artist)
				}

				if *album != "" {
					audioFile.SetAlbum(*album)
				}
			}
		}
		return nil
	})
}

// HasAudioExtension returns true if has audio extension
func HasAudioExtension(path string) bool {
	supportedTypes := [...]string{
		".mp3",
		".m4a",
		".ogg",
		".flac",
		".wav",
	}

	pathExtension := filepath.Ext(path)
	for _, supportedExt := range supportedTypes {
		if supportedExt == pathExtension {
			return true
		}
	}

	return false
}
