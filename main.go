package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/corykitchens/wavvy/pkg/kcpr"
)

func main() {
	jsonFlag := flag.Bool("json", false, "Return the current output as json")
	flag.Parse()
	track, err := kcpr.GetCurrentTrack()
	if err != nil {
		log.Fatal("Error retrieving current track")
	}
	if *jsonFlag == true {
		fmt.Println(track.Json())
		return
	}
	fmt.Println(track)
}
