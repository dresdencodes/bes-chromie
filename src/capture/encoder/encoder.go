package encoder

import (
	"io"
	"bytes"
	//"image/png"
	"os/exec"
)

type Encoder struct {
	Command 		*exec.Cmd
	Stdin 			io.WriteCloser
	Stdout			io.ReadCloser
}

func New() (*Encoder, error) {

	// defs
	e := &Encoder{}
	var err error

	// FFmpeg command: read PNG frames from stdin and make a video
	e.Command = exec.Command("ffmpeg",
		"-y",                // overwrite output
		"-f", "image2pipe",  // stream images
		"-vcodec", "png",    // tell ffmpeg they're PNGs
		"-framerate", "30",  // set input framerate
		"-i", "pipe:0",      // read from stdin
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		"out.mp4",
	)

	// start pipe
	e.Stdin, err = e.Command.StdinPipe()
	if err != nil {
		return e, err
	}

	// start command
	err = e.Command.Start() 
	return e, err
	
}

func (e *Encoder) Finish() error {

	//
	// stdin close
	//
	e.Stdin.Close() 
	err := e.Command.Wait()
	return err

}

func (e *Encoder) AddPNG(buf *bytes.Buffer) error {
	
	//
	// write
	//
	_, err := e.Stdin.Write(buf.Bytes())
	return err
	
}