package main

import (
	"io"
	"os"
	"strings"
)

const ky13 = uint8(13)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	if _, err := rot13.r.Read(p); err != nil {
		return 0, err
	}

	for iI, yVal := range p {
		switch {
		case (yVal >= 64) && (yVal <= 77):
			fallthrough
		case (yVal >= 97) && (yVal <= 109):
			p[iI] += ky13
		case (yVal >= 78) && (yVal <= 90):
			fallthrough
		case (yVal >= 110) && (yVal <= 122):
			p[iI] -= ky13
		}
	}

	return len(p), nil
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
