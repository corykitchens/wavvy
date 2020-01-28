package kcpr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const xmlResp string = `<?xml version="1.0" encoding="UTF-8"?>
<playlist>
    <DCSoutputVersion>2</DCSoutputVersion>
    <stationCallSign>KCPR1</stationCallSign>
    <programType>PGM</programType>
    <mediaType>AUD</mediaType>
    <title>Rugged Country</title>
    <artist>Japanese Breakfast</artist>
    <album>Psychopomp</album>
    <cover>https://cdnrf.securenetsystems.net/file_radio/album_art/q/1/5/51qbh01s04L.jpg</cover>
    <duration>0</duration>
    <campaignId></campaignId>
    <fileId></fileId>
    <programStartTS>28 Jan 2020 21:29:23</programStartTS>
    <adBlockPos></adBlockPos>
</playlist>`

type mockClient struct {
}

func (m mockClient) Get(url string) (*http.Response, error) {
	h := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(xmlResp)),
	}

	return h, nil
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %s, Want %s", got, want)
	}

}
func TestTrackFields(t *testing.T) {
	for _, tt := range trackFieldTestCases {
		track := Track{Artist: tt.Artist, Title: tt.Title, Album: tt.Album}
		assertEqual(t, track.Artist, tt.Artist)
		assertEqual(t, track.Album, tt.Album)
		assertEqual(t, track.Title, tt.Title)
	}
}

func TestTrackString(t *testing.T) {
	for _, tt := range trackStringTestCases {
		track := Track{Artist: tt.track.Artist, Title: tt.track.Title, Album: tt.track.Album}
		buf := bytes.Buffer{}
		fmt.Fprintf(&buf, "%s", track)
		got := buf.String()
		assertEqual(t, got, tt.want)
	}
}

func TestTrackJson(t *testing.T) {
	for _, tt := range trackJsonTestCases {
		track := Track{Artist: tt.track.Artist, Title: tt.track.Title, Album: tt.track.Album}
		got, err := track.Json()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		assertEqual(t, got, tt.want)
	}
}

func TestGetXmlResponse(t *testing.T) {
	m := mockClient{}
	url := "http://hello.com"
	resp, err := getXmlResponse(m, url)
	if err != nil {
		t.Errorf("Got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	bodyString := string(data)

	if err != nil {
		t.Errorf("Got %v", err)
	}
	if bodyString != xmlResp {
		t.Errorf("Got %s, Want %s", bodyString, xmlResp)
	}
}

func TestConvertResponseToTrack(t *testing.T) {
	want := Track{Artist: "Japanese Breakfast", Title: "Rugged Country", Album: "Psychopomp"}
	got, err := convertXmlResponseToTrack(xmlResp)
	if err != nil {
		t.Errorf("Received error %v, Wanted nil", err)
	}
	if got != want {
		t.Errorf("Got %s, Want %s", got, want)
	}
}
