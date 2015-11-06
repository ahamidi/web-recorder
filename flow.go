package recorder

import (
	"image/png"
	"io"
	"log"
)

func Flow(in io.ReadCloser, out io.WriteCloser, fps int) {
	// TODO:
	// - Hold onto last "frame"
	// - If "tick" elapses without receiving new frame, send old

	for in != nil {
		img, err := png.Decode(in)
		if err != nil {
			out.Close()
			//log.Println("Error!:", err)
			return
		}

		err = png.Encode(out, img)
		if err != nil {
			log.Println("Error!!:", err)
			return
		}
	}

	out.Close()
}
