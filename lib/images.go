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

type imageMap = map[string]*image.Image

// Images : simple map of images
var Images = make(imageMap)

var mu sync.Mutex
var wd, _ = os.Getwd()

func init() {

	for _, imagePath := range imagePaths {
		loadFile(imagePath, Images, &mu)
	}
}

func loadFile(imagePath string, images imageMap, mu *sync.Mutex) {
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
	images[imagePath] = &img
	mu.Unlock()
}
