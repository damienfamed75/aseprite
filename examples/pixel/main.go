package main

import (
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/damienfamed75/aseprite"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	// Create a new pixel window.
	cfg := pixelgl.WindowConfig{
		Title:     "Demo",
		Bounds:    pixel.R(0, 0, screenWidth, screenHeight),
		Resizable: true,
	}
	win, _ := pixelgl.NewWindow(cfg)

	// Open the player spritesheet's aseprite JSON file.
	ase, _ := aseprite.Open("player.json")
	ase.Play("left")

	// Open the player image
	pic, err := loadPicture("player.png")
	if err != nil {
		panic(err)
	}

	// Creates a pixel Matrix Drawing object.
	imd := imdraw.New(pic)
	sprite := pixel.NewSprite(pic, pic.Bounds())
	last := time.Now()
	for !win.Closed() {
		// Calculate delta time. I'm sure there's a better option on how to
		// create this, because everytime I run this the animation stutters.
		dt := float32(time.Since(last).Seconds() / 2)
		last = time.Now()

		ase.Update(dt)
		bounds := ase.FrameBoundaries().Rectangle()

		imd.Clear()

		sprite.Set(pic, pixel.R(
			float64(bounds.Min.X), float64(bounds.Min.Y),
			float64(bounds.Max.X), float64(bounds.Max.Y)),
		)
		sprite.Draw(
			win, pixel.IM.Scaled(pixel.ZV, 4.0).Moved(
				win.Bounds().Center(),
			),
		)
		imd.Draw(win)

		win.Update()
		win.Clear(colornames.Skyblue)
	}
}
