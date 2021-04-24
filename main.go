package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	diff := max - min
	if diff == 0 {
		return rand.Intn(max)
	}
	return rand.Intn(diff) + min
}

func generateWidthAndHeight(mw, mh int) (width, height int) {
	width = generateRandomNumber(mw, mh)
	height = generateRandomNumber(mw, mh)
	return
}

func generateRGB() (r, g, b float64) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 255
	r = float64(rand.Intn(max-min) + min)
	g = float64(rand.Intn(max-min) + min)
	b = float64(rand.Intn(max-min) + min)
	return
}

func generate(file string, iterations int) {
	gl, err := glitch.NewGlitch("./" + file)
	if err != nil {
		panic("could ont create new glitch")
	}
	width, height := gl.Bounds.Max.X, gl.Bounds.Max.Y

	// Directory

	directory := fmt.Sprintf("glich_%s", file)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, os.ModePerm)
	}

	for i := 0; i < iterations; i++ {
		rand.Seed(time.Now().UnixNano())
		gl.Seed(rand.Int63())

		gl.Copy()

		filename := fmt.Sprintf("./glitch_%d_%s", i+1, file)
		fmt.Printf("[%d/%d] Generating %s\n", i+1, iterations, filename)

		// Lineas verticales

		//{
		//	w, h := generateWidthAndHeight(width, height)
		//	gl.VerticalTransposeInput(w/10, h/10, true)
		//}

		// Horizontales

		//{
		//	w, h := generateWidthAndHeight(width, height)
		//	gl.TransposeInput(h/10, w/10, true)
		//}

		{
			w, h := generateWidthAndHeight(width, height)
			if w < h {
				gl.ChannelShiftLeft()
				gl.HalfLifeLeft(w, width)
			} else {
				gl.ChannelShiftRight()
				gl.HalfLifeRight(h, width)
			}
		}

		//{
		//	w, h := generateWidthAndHeight(width, height)
		//	if w < h {
		//		gl.BlueBoost()
		//	} else {
		//		gl.GreenBoost()
		//	}
		//}

		//{
		//	r, g, b := generateRGB()
		//	gl.Noise(r, g, b, 0.07)
		//}

		gl.GhostStreach()
		gl.RedBoost()

		f, err := os.Create(filename)
		if err != nil {
			panic("could not create new file")
		}
		gl.Write(f)

		// move
		err = os.Rename(filename, fmt.Sprintf("%s/%s", directory, filename))
		if err != nil {
			log.Fatal(err)
		}
	}
}
