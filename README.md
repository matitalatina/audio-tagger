# Audio Tagger

Simple app to change artist and album for audio files in bulk inside a folder.

## Getting Started

### Dependencies

- `go get github.com/mikkyang/id3-go` -> id3 Tag Library.
- `go get github.com/h2non/filetype` -> FileType check.
- `go get github.com/stretchr/testify` -> Testing.

Or simply `go get -t -v ./...` to fetch them all.

### CLI

`go run main.go -folder FOLDER/THAT/STORES/AUDIO/FILES -album "Album Name" -artist "Artist Name"`

## Test

`go test`