package main

import (
	"bytes"
	"fmt"
	"github.com/jpoz/glitch"
	"image"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"time"
)

func createGlitch(filename string) *glitch.Glitch {
	gl, err := glitch.NewGlitch(filename)
	if err != nil {
		log.Fatalf("could not create new glitch: %v", err)
	}
	return gl
}

func start(filename string, iterations uint) {
	gl := createGlitch(filename)
	generate(gl, filename, iterations)
}

func generate(gl *glitch.Glitch, filename string, iterations uint) {
	file, ext, directory := getFileDetails(filename)

	// Directory

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.MkdirAll(directory, os.ModePerm)
	}

	files := make([]string, 0)

	for i := uint(0); i < iterations; i++ {
		rand.Seed(time.Now().UnixNano())
		gl.Seed(rand.Int63())

		gl.Copy()
		glitchIt(gl)

		filename := createFileName(file, ext, i)
		fmt.Printf("[%d/%d] Generating %s\n", i+1, iterations, filename)
		f, err := os.Create(filename)
		if err != nil {
			panic("could not create new file")
		}
		gl.Write(f)

		files = append(files, filename)
		// move
		err = os.Rename(filename, fmt.Sprintf("%s/%s", directory, filename))
		if err != nil {
			log.Fatal(err)
		}
	}

	createAnimation(files, directory, file)
}

func createAnimation(files []string, directory, filename string) {
	fmt.Println("Generating GIF")
	outGif := &gif.GIF{}
	total := len(files)

	for i, name := range files {
		fmt.Printf("[%d/%d] Processing GIF\n", i+1, total)
		input := fmt.Sprintf("%s/%s", directory, name)

		f, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}

		imageData, _, err := image.Decode(f)
		if err != nil {
			log.Println("Error al decorde archivo")
			log.Fatal(err)
		}

		buf := bytes.Buffer{}

		if err = gif.Encode(&buf, imageData, nil); err != nil {
			log.Println("Error al encode archivo")
			log.Fatal(err)
		}

		inGif, err := gif.Decode(&buf)
		if err != nil {
			log.Println("Erro en gif decode")
			log.Fatal(err)
		}
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	output := fmt.Sprintf("FINAL_%s.gif", createGifName(filename))

	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	gif.EncodeAll(f, outGif)

	err = os.Rename(output, fmt.Sprintf("%s/%s", directory, output))
	if err != nil {
		log.Fatal(err)
	}

}
