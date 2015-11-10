package main

import (
	"bytes"
	"flag"
	"io"
	"log"

	wr "github.com/ahamidi/web-recorder"
)

var url = flag.String("url", "http://ahamidi.com", "URL to Record.")
var duration = flag.Int("duration", 5, "Duration of recording (in seconds).")

func main() {
	log.Println("Launching Recorder...")
	flag.Parse()

	rec, err := wr.NewRecorder(*url)
	if err != nil {
		panic(err)
	}
	rec.Duration = *duration
	p, err := wr.NewPhantom(rec)
	if err != nil {
		panic(err)
	}
	ff, err := wr.NewFfmpeg(rec)
	if err != nil {
		panic(err)
	}

	reader, writer := io.Pipe()
	p.Cmd.Stdout = writer

	r2, w2 := io.Pipe()
	ff.Cmd.Stdin = r2

	go wr.Flow(reader, w2, 10)

	var buff bytes.Buffer
	p.Cmd.Stderr = &buff

	ff.Cmd.Start()
	p.Cmd.Run()
	writer.Close()
	ff.Cmd.Wait()

	if buff.Len() != 0 {
		log.Println("Errout:", buff.String())
	}

}
