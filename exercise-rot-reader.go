package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		caseInsensetiveLetter := strings.ToLower(string(b[i]))
		if caseInsensetiveLetter >= "a" && caseInsensetiveLetter <= "m" {
			b[i] = byte(int(b[i]) + 13)
		} else if caseInsensetiveLetter > "m" && caseInsensetiveLetter <= "z" {
			b[i] = byte(int(b[i]) - 13)
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
