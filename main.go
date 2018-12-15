package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	id3 "github.com/bogem/id3v2"

	"github.com/h2non/filetype"
)

type TagRequest struct {
	album, artist string
}

func main() {
	folder := flag.String("folder", "", "[Required] Folder where the audio files are stored")
	album := flag.String("album", "", "Album to set")
	artist := flag.String("artist", "", "Artist to set")
	flag.Parse()

	if flag.NFlag() < 2 || *folder == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	tagRequest := TagRequest{*album, *artist}

	filepath.Walk(*folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return doOnlyAudio(path, func() {
			fmt.Println(path)
			audioFile, _ := id3.Open(path, id3.Options{Parse: true})
			defer audioFile.Close()
			setTags(audioFile, tagRequest)
		})
	})
}

func doOnlyAudio(path string, f func()) error {
	if HasAudioExtension(path) {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if filetype.IsAudio(buf) {
			f()
		}
	}
	return nil
}

func setTags(audio *id3.Tag, tagRequest TagRequest) {
	if tagRequest.artist != "" {
		audio.SetArtist(tagRequest.artist)
	}

	if tagRequest.album != "" {
		audio.SetAlbum(tagRequest.album)
	}

	audio.Save()
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
