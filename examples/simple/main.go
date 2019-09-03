package main

import (
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/damienfamed75/aseprite"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	count int

	old time.Time
	new = time.Now()

	player      *ebiten.Image
	playerSheet *aseprite.File
)

func main() {
	playerImage, err := os.Open("player.png")
	if err != nil {
		log.Fatalln("player open:", err)
	}

	img, _, err := image.Decode(playerImage)
	if err != nil {
		log.Fatalf("player decode: err[%v]\n", err)
	}

	player, err = ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if err != nil {
		log.Fatalln("player ebiten new image:", err)
	}

	animSheet, err := os.Open("player.json")
	if err != nil {
		log.Fatalln(err)
	}

	bytes, err := ioutil.ReadAll(animSheet)
	if err != nil {
		log.Fatalln(err)
	}

	playerSheet, err = aseprite.NewFile(bytes)
	if err != nil {
		log.Fatalln(err)
	}

	playerSheet.Play("down")

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Demo"); err != nil {
		log.Fatalln(err)
	}
}

func update(screen *ebiten.Image) error {
	count++
	old = new
	new = time.Now()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	dt := (float32(new.Sub(old).Milliseconds()) / 1000) / 2
	playerSheet.Update(dt)

	op := &ebiten.DrawImageOptions{}
	bounds := playerSheet.FrameBoundaries()
	op.GeoM.Translate(-float64(bounds.Width)/2, -float64(bounds.Height)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	screen.DrawImage(
		player.SubImage(bounds.Rectangle()).(*ebiten.Image), op,
	)

	return nil
}
