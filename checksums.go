package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	pause := false
	var filename string
	for _, arg := range os.Args[1:] {
		if arg == "-p" {
			pause = true
		} else {
			filename = arg
		}
	}

	calc(filename)

	if pause {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Press ENTER key to exit...")
		reader.ReadString('\n')
	}
}

type fanWriter struct {
	outputs []io.Writer
}

func (w *fanWriter) add(ww io.Writer) {
	w.outputs = append(w.outputs, ww)
}

func (w *fanWriter) Write(p []byte) (n int, err error) {
	for _, out := range w.outputs {
		out.Write(p)
	}
	return len(p), nil
}

func calc(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can not open file: ", filename)
		return
	}
	defer f.Close()

	writers := &fanWriter{}

	hMD5 := md5.New()
	writers.add(hMD5)

	hSHA1 := sha1.New()
	writers.add(hSHA1)

	hSHA256 := sha256.New()
	writers.add(hSHA256)

	if _, err := io.Copy(writers, f); err != nil {
		fmt.Println("Can not read file: ", filename)
		return
	}

	fmt.Printf("MD5    : %032x\n", hMD5.Sum(nil))
	fmt.Printf("SHA1   : %040x\n", hSHA1.Sum(nil))
	fmt.Printf("SHA256 : %064x\n", hSHA256.Sum(nil))
}
