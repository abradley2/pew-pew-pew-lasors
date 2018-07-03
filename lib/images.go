package lib

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"sync"
)

var imagePaths = []string{
	"/assets/lazor-green.png",
	"/assets/lazor-red.png",
	"/assets/tie-smol.png",
	"/assets/xwing-smol.png",
}

type imageMap = map[string]image.Image

// Images : a map of image paths to their byte slices
var Images = make(map[string]image.Image)

var mu sync.Mutex
var wd, _ = os.Getwd()

func init() {
	fileLoadChannel := make(chan error)

	for _, imagePath := range imagePaths {
		go loadFile(imagePath, Images, &mu, fileLoadChannel)
	}

	<-fileLoadChannel
}

func loadFile(imagePath string, images imageMap, mu *sync.Mutex, channel chan<- error) {
	data, err := os.Open(filepath.Join(wd, imagePath))
	if err != nil {
		panic(fmt.Sprintf("issue opening image %s\n", imagePath))
	}
	img, imageDecodeErr := png.Decode(data)
	data.Close()

	if imageDecodeErr != nil {
		panic(fmt.Sprintf("issue decoding image %s\n", imagePath))
	}

	mu.Lock()
	images[imagePath] = img
	mu.Unlock()

	channel <- err
}
