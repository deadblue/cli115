package core

import "io"

func QuietlyClose(c io.Closer) {
	if c != nil {
		_ = c.Close()
	}
}
