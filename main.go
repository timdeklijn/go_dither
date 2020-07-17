package main

import (
	"flag"
	"fmt"

	"github.com/timdeklijn/go_dither/fsdither"
)

func main() {
	pathFlag := flag.String("file", "images/cat.jpg", "File to Dither")
	flag.Parse()
	filePath := *pathFlag
	fmt.Println(filePath)
	fsdither.Dither(filePath)
}
