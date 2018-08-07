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
	writers []io.Writer
}

func (fw *fanWriter) add(w io.Writer) {
	fw.writers = append(fw.writers, w)
}

func (fw *fanWriter) Write(p []byte) (n int, err error) {
	for _, w := range fw.writers {
		w.Write(p)
	}
	return len(p), nil
}

func calc(filename string) {
	reader, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can not open file: ", filename)
		return
	}
	defer reader.Close()

	fw := &fanWriter{}

	hMD5 := md5.New()
	fw.add(hMD5)

	hSHA1 := sha1.New()
	fw.add(hSHA1)

	hSHA256 := sha256.New()
	fw.add(hSHA256)

	if _, err := io.Copy(fw, reader); err != nil {
		fmt.Println("Can not read file: ", filename)
		return
	}

	fmt.Printf("MD5    : %032x\n", hMD5.Sum(nil))
	fmt.Printf("SHA1   : %040x\n", hSHA1.Sum(nil))
	fmt.Printf("SHA256 : %064x\n", hSHA256.Sum(nil))
}
