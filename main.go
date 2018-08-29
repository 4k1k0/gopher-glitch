package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jpoz/glitch"
)

func main() {

	file_name := flag.String("file", "", "the name of your image")
	iterations := flag.Int("n", 5, "the number of files")

	flag.Parse()

	if *file_name == "" {
		fmt.Println("Select an image")
		os.Exit(1)
	}

	generate(*file_name, *iterations)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generate(file string, iterations int) {
	gl, err := glitch.NewGlitch("./" + file)
	check(err)
	gl.Copy()

	for i := 0; i < iterations; i++ {
		n := strconv.Itoa(i + 1)
		gl.VerticalTransposeInput(i+1*50, 100, true)
		gl.ChannelShiftLeft()

		newFile := fmt.Sprintf("./glitch_" + n + "_" + file)
		f, err := os.Create(newFile)
		check(err)
		gl.Write(f)
	}
}
