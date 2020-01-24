package kcpr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

//A Track groups the currently
//played track's artist, song title, and album
type Track struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Album  string `json:"album"`
}

//String method returns the following formatted string
//Artist: <artist_name> Title: <song_title> Album: <album_title>
func (t Track) String() string {
	return fmt.Sprintf("Artist: %s, Title: %s, Album: %s", t.Artist, t.Title, t.Album)
}

//Json method returns the track as json
//{artist: <artist_name>, title: <song_title>, album: <album_title> }
func (t Track) Json() (string, error) {
	b := bytes.Buffer{}
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(t)
	if err != nil {
		return "", err
	}
	parsedString := strings.TrimSuffix(b.String(), "\n")
	return parsedString, nil
}
