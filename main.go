package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("filename", "", "the name of your image")
	iterations := flag.Uint("times", 2, "the number of frames")

	flag.Parse()

	if *filename == "" {
		fmt.Println("Select an image")
		os.Exit(1)
	}

	start(*filename, *iterations)
}
