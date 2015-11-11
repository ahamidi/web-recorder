package recorder

import (
	"bytes"
	"io"
	"log"
	"net/url"
	"os/exec"
)

type ClipArea struct {
	Top    int
	Left   int
	Width  int
	Height int
}

type ViewPort struct {
	Width  int
	Height int
}

type Recorder struct {
	URL          *url.URL
	Duration     int
	Framerate    int
	OutputFile   string
	InputClip    *ClipArea
	ViewportSize *ViewPort
	Recorder     *Command
	Transcoder   *Command
}

type Command struct {
	Cmd    *exec.Cmd
	In     io.WriteCloser
	Out    io.ReadCloser
	Errout io.ReadCloser
}

func NewRecorder(targetURL string, duration int, output string) (*Recorder, error) {
	log.Println("New Recorder created.")

	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	if duration == 0 {
		duration = 5
	}

	if output == "" {
		output = "output.mp4"
	}

	r := &Recorder{
		URL:        u,
		Duration:   duration,
		Framerate:  25,
		OutputFile: output,
	}

	r.Recorder, err = NewPhantom(r)
	if err != nil {
		panic(err)
	}
	r.Transcoder, err = NewFfmpeg(r)
	if err != nil {
		panic(err)
	}

	return r, nil
}

func (r *Recorder) Start() {

	go Flow(r.Recorder.Out, r.Transcoder.In, 10)

	var buff bytes.Buffer
	r.Recorder.Cmd.Stderr = &buff

	r.Transcoder.Cmd.Start()
	r.Recorder.Cmd.Run()
	r.Transcoder.Cmd.Wait()

	if buff.Len() != 0 {
		log.Println("Errout:", buff.String())
	}

}
