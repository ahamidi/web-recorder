package recorder

import (
	"image"
	"image/png"
	"io"
	"log"
	"time"
)

func Flow(in io.ReadCloser, out io.WriteCloser, fps int) {
	// TODO:
	// - Hold onto last "frame"
	// - If "tick" elapses without receiving new frame, send old

	td := time.Duration(1000 / fps)
	ticker := time.NewTicker(time.Millisecond * td)

	var img image.Image
	var err error
	done := make(chan bool, 1)

	go func() {
		for in != nil {
			img, err = png.Decode(in)
			if err != nil {
				//log.Println("Error!:", err)
				break
			}
		}
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			if img != nil {
				err := png.Encode(out, img)
				if err != nil {
					log.Println("Error!!:", err)
				}
			}

		case <-done:
			out.Close()
		}

	}
}
