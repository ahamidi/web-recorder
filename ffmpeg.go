package recorder

import (
	"os"
	"os/exec"
	//"strconv"
)

var FfmpegPath = "ffmpeg"

func NewFfmpeg(r *Recorder) (*Command, error) {

	// Check if custom phantomjs bin path
	if path := os.Getenv("FFMPEG_BIN"); path != "" {
		FfmpegPath = path
	}

	cmd := exec.Command(FfmpegPath, "-r", "10", "-y", "-c:v", "png", "-f", "image2pipe", "-i", "-", "-c:v", "libx264", "-pix_fmt", "yuv420p", "-movflags", "+faststart", "/tmp"+r.OutputFile)

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

	ffm := &Command{
		Cmd:    cmd,
		In:     inPipe,
		Out:    outPipe,
		Errout: errPipe,
	}

	return ffm, nil

}
