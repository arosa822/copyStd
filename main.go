package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"io"
	"os"
)

func test(s string) string {
	return s
}

func copyToClipboard(s string) {
	clipboard.WriteAll(s)
	fmt.Println("Copying output to Clipboard...\n")
	return
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}

}

func main() {
	//s := test("this is a test")
	message := captureStdout(scanner)
	copyToClipboard(message)
}
