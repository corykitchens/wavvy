package kcpr

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const kcprUrl = "https://streamdb8web.securenetsystems.net/player_status_update/KCPR1.xml"

type Requester interface {
	Get(string) (*http.Response, error)
}

//A Track groups the currently
//played track's artist, song title, and album
type Track struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Album  string `json:"album"`
}

type XMLResponse struct {
	XMLName xml.Name `xml:"playlist"`
	Title   string   `xml:"title"`
	Artist  string   `xml:"artist"`
	Album   string   `xml:"album"`
	Cover   string   `xml:"cover"`
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

func getXmlResponse(r Requester, url string) (*http.Response, error) {
	resp, err := r.Get(url)
	if err != nil {
		return nil, errors.New("Error sending request")
	}
	return resp, nil
}

func convertXmlResponseToTrack(data []byte) (*Track, error) {
	var xmlResp XMLResponse
	xml.Unmarshal(data, &xmlResp)
	t := &Track{
		Artist: xmlResp.Artist,
		Title:  xmlResp.Title,
		Album:  xmlResp.Album,
	}
	return t, nil
}

func GetCurrentTrack() (*Track, error) {
	h := &http.Client{}
	resp, err := getXmlResponse(h, kcprUrl)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	track, err := convertXmlResponseToTrack(data)
	if err != nil {
		return nil, err
	}
	return track, nil
}
