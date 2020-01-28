package kcpr

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
