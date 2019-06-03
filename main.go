package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

const (
	screenWidth  = 240
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

const (
	tileSize = 16
	tileXNum = 25
)

var (
	tilesImage 	*ebiten.Image
	count       = 0
	runnerImage *ebiten.Image
)

func init() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

var (
	layers = [][]int{
		{
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
		},
		{
			  0,   0,   0,   0,   0, 303,   0,   0,   0,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,  26,  27,  28,  29,  30,  31,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,  51,  52,  53,  54,  55,  56,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,  76,  77,  78,  79,  80,  81,   0,   0,   0,   0,
			  0,   0,   0,   0,   0, 101, 102, 103, 104, 105, 106,   0,   0,   0,   0,
			  0,   0,   0,   0,   0, 126, 127, 128, 129, 128, 131,   0,   0,   0,   0,
			  0,   0,   0,   0,   0, 303, 303, 245, 242, 303, 303,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
			  0,   0,   0,   0,   0,   0,   0, 245, 242,   0,   0,   0,   0,   0,   0,
		},
	}
)

func update(screen *ebiten.Image) error {
	count++

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	const xNum = screenWidth / tileSize
	for _, l := range layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t % tileXNum) * tileSize
			sy := (t / tileXNum) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	return nil
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "GooeyBlob"); err != nil {
		log.Fatal(err)
	}
}