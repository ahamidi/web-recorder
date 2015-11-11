package recorder

import (
	"os"
	"os/exec"
	"strconv"
)

var PhantomJSBinPath = "phantomjs"

func NewPhantom(r *Recorder) (*Command, error) {

	// Check if custom phantomjs bin path
	if path := os.Getenv("PHANTOMJS_BIN"); path != "" {
		PhantomJSBinPath = path
	}

	cmd := exec.Command(PhantomJSBinPath, "capture_video.js", r.URL.String(), strconv.Itoa(r.Duration))

	// Map various pipes
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	pjs := &Command{
		Cmd:    cmd,
		In:     inPipe,
		Out:    outPipe,
		Errout: errPipe,
	}

	return pjs, nil
}
