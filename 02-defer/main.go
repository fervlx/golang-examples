package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := write("logs.txt", "01/78.p import err, line 45637"); err != nil {
		log.Fatal("failed to write file", err)
	}

	fmt.Println("write: success.")

	if err := copy("logs.txt", "log2.txt"); err != nil {
		log.Fatal("failed to copy file", err)
	}

	fmt.Println("copy: success.")
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.WriteString(file, text)

	if err != nil {
		return err
	}

	return file.Close()
}

func copy(source string, destination string) error {
	src, err := os.Open(source)

	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(destination)

	if err != nil {
		return err
	}

	defer dst.Close()

	size, err := io.Copy(dst, src)

	if err != nil {
		return err
	}

	fmt.Printf("Copied %d bytes from %s to %s\n", size, source, destination)

	if err := src.Close(); err != nil {
		return err
	}

	return dst.Close()
}
