package main

// For the sake of brevity this example ignores error handling.

import (
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/damienfamed75/aseprite"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240

	renderScale = 4.0
)

var (
	old time.Time
	new = time.Now()

	player      *ebiten.Image
	playerSheet *aseprite.File
)

func main() {
	// Open player image in ebiten and load it into memory.
	playerImage, _ := os.Open("player.png")
	img, _, _ := image.Decode(playerImage)
	player, _ = ebiten.NewImageFromImage(img, ebiten.FilterNearest)

	// Open the player spritesheet's aseprite JSON file and play a default animation.
	playerSheet, _ = aseprite.Open("player.json")
	playerSheet.Play("right")

	// Run the app with a game loop.
	ebiten.Run(update, screenWidth, screenHeight, 2, "Demo")
}

func update(screen *ebiten.Image) error {
	// Keep track of the delta time however you wish.
	// This is how I'm doing it for the example. It may not be the best solution.
	old = new
	new = time.Now()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Calculate delta time.
	dt := (float32(new.Sub(old).Milliseconds()) / 1000) / 2
	// Get the current frame's bounding box.
	bounds := playerSheet.FrameBoundaries()
	// Update the spritesheet.
	playerSheet.Update(dt)

	// Create options that will center the sub image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(renderScale, renderScale)
	op.GeoM.Translate(
		-float64(bounds.Width*renderScale)/2,
		-float64(bounds.Height*renderScale)/2,
	)
	op.GeoM.Translate(
		screenWidth/2,
		screenHeight/2,
	)

	// Draw the player.
	screen.DrawImage(
		// Create a sub image with the bounding box's Rectangle.
		player.SubImage(bounds.Rectangle()).(*ebiten.Image), op,
	)

	return nil
}
