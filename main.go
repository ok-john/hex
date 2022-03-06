package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	if len(os.Args) <= 1 {
		fmt.Println("[USAGE]\n\nencoding:\n\thex [stuff]\ndecoding:\n\thex -d [stuff]")
		return
	}
	if os.Args[1] == "-d" {
		input = strings.Join(os.Args[2:], "")
		os.Stdout.Write(<-DecodeHex([]byte(input)))
	} else {
		input = strings.Join(os.Args[1:], "")
		os.Stdout.Write(<-EncodeHex([]byte(input)))
	}
	fmt.Println("")
}
