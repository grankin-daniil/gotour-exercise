package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func test() int {
	return 0
}

func (reader MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
