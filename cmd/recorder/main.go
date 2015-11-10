package main

import (
	"flag"
	"log"

	wr "github.com/ahamidi/web-recorder"
)

var url = flag.String("url", "http://ahamidi.com", "URL to Record.")
var duration = flag.Int("duration", 5, "Duration of recording (in seconds).")

func main() {
	log.Println("Launching Recorder...")
	flag.Parse()

	rec, err := wr.NewRecorder(*url, *duration)
	if err != nil {
		panic(err)
	}

	rec.Start()
}
