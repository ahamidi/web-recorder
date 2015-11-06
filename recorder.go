package recorder

import (
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
}

type Command struct {
	Cmd    *exec.Cmd
	In     io.WriteCloser
	Out    io.ReadCloser
	Errout io.ReadCloser
}

func NewRecorder(targetURL string) (*Recorder, error) {
	log.Println("New Recorder created.")

	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	r := &Recorder{
		URL:        u,
		Duration:   5,
		Framerate:  25,
		OutputFile: "output.mp4",
	}

	return r, nil
}
