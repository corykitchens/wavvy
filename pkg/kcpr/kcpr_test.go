package kcpr

import "testing"

var trackTestCases = []struct {
	Artist string
	Title  string
	Album  string
}{
	{"King Gizzard and the Lizard Wizard", "Cyboogie", "Fishing for Fishes"},
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %s, Want %s", got, want)
	}

}
func TestTrack(t *testing.T) {
	for _, tt := range trackTestCases {
		track := Track{Artist: tt.Artist, Title: tt.Title, Album: tt.Album}
		assertEqual(t, track.Artist, tt.Artist)
		assertEqual(t, track.Album, tt.Album)
		assertEqual(t, track.Title, tt.Title)
	}
}
