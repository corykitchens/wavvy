package kcpr

import (
	"bytes"
	"fmt"
	"testing"
)

var trackFieldTestCases = []struct {
	Artist string
	Title  string
	Album  string
}{
	{"King Gizzard and the Lizard Wizard", "Cyboogie", "Fishing for Fishes"},
}

var trackStringTestCases = []struct {
	track Track
	want  string
}{
	{
		track: Track{"King Gizzard and the Lizard Wizard", "Cyboogie", "Fishing for Fishes"},
		want:  "Artist: King Gizzard and the Lizard Wizard, Title: Cyboogie, Album: Fishing for Fishes",
	},
}

var trackJsonTestCases = []struct {
	track Track
	want  string
}{
	{
		track: Track{"King Gizzard and the Lizard Wizard", "Cyboogie", "Fishing for Fishes"},
		want:  `{"artist":"King Gizzard and the Lizard Wizard","title":"Cyboogie","album":"Fishing for Fishes"}`,
	},
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
