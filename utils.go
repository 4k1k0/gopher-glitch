package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	diff := max - min
	if diff <= 0 {
		return rand.Intn(max)
	}
	return rand.Intn(diff) + min
}

func generateWidthAndHeight(mw, mh int) (width, height int) {
	width = generateRandomNumber(0, mw)
	height = generateRandomNumber(0, mh)
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

func getFileDetails(path string) (filename, ext, directory string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open %q\n", path)
	}
	defer f.Close()

	ext = filepath.Ext(path)
	filename = strings.TrimSuffix(filepath.Base(path), ext)
	filename = strings.Replace(filename, " ", "_", -1)
	directory = fmt.Sprintf("%s/%s", config.Path, filename)

	return
}

func createFileName(file, ext string, index uint) string {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	unixTime := currentTime.Unix()
	return fmt.Sprintf("%s_%s_%d_%d%s", file, date, index+1, unixTime, ext)
}

func createGifName(file string) string {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	unixTime := currentTime.Unix()
	return fmt.Sprintf("%s_%s_%d", file, date, unixTime)
}