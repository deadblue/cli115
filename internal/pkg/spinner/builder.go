package spinner

import (
	"io"
	"os"
	"time"
)

const (
	FpsDefault = 10
)

var (
	FrameSetDefault = []string{
		"\\", "|", "/", "-",
	}
)

type Builder struct {
	out      io.Writer
	fps      int
	frames   []string
	prefix   string
	suffix   string
	complete string
}

func (b *Builder) Fps(fps int) *Builder {
	if fps > 0 {
		b.fps = fps
	}
	return b
}

func (b *Builder) Frames(frames []string) *Builder {
	b.frames = frames
	return b
}

func (b *Builder) Prefix(text string) *Builder {
	b.prefix = text
	return b
}

func (b *Builder) Suffix(text string) *Builder {
	b.suffix = text
	return b
}

func (b *Builder) Complete(text string) *Builder {
	b.complete = text
	return b
}

func (b *Builder) Build() Spinner {
	if b.fps == 0 {
		b.fps = FpsDefault
	}
	interval := time.Second / time.Duration(b.fps)
	return &implSpinner{
		out:      b.out,
		interval: interval,

		frames:   b.frames,
		prefix:   b.prefix,
		suffix:   b.suffix,
		complete: b.complete,
	}
}

func NewBuilder() *Builder {
	return &Builder{
		out: os.Stdout,
		fps: FpsDefault,

		frames:   FrameSetDefault,
		prefix:   "",
		suffix:   "",
		complete: "",
	}
}
