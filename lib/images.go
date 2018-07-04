package lib

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

var imagePaths = []string{
	"/assets/lazor-green.png",
	"/assets/lazor-red.png",
	"/assets/tie-smol.png",
	"/assets/xwing-smol.png",
}

type imageMap = map[string]ebiten.Image

// Images : a map of image paths to their byte slices
var Images = make(map[string]ebiten.Image)

var mu sync.Mutex
var wd, _ = os.Getwd()

// EbitenImage : the main stage
var EbitenImage *ebiten.Image

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

	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := origEbitenImage.Size()
	EbitenImage, _ = ebiten.NewImage(w, h, ebiten.FilterNearest)

	op := &ebiten.DrawImageOptions{}
	EbitenImage.DrawImage(origEbitenImage, op)

	mu.Lock()
	images[imagePath] = *origEbitenImage
	mu.Unlock()

	channel <- err
}
