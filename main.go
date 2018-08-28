package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jpoz/glitch"
)

func main() {

	check_errors()

	file_name := os.Args[1]
	_, error := strconv.Atoi(os.Args[2])
	if error != nil {
		fmt.Println("Not a number")
		os.Exit(1)
	}

	n, _ := strconv.Atoi(os.Args[2])

	generate(file_name, n)
}

func check_errors() {
	if len(os.Args) < 2 {
		fmt.Println("Add an image")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("Add a number")
		os.Exit(1)
	}
	if len(os.Args) != 3 {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}
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
