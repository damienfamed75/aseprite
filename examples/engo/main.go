package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/damienfamed75/aseprite"
)

const (
	screenWidth  = 320
	screenHeight = 240

	entityScale = 4.0
)

func main() {
	opts := engo.RunOptions{
		Title:  "Demo",
		Width:  screenWidth,
		Height: screenHeight,
	}

	engo.Run(opts, &Demo{})
}

// Demo is a scene used to demonstrate animating a spritesheet using aseprite.
type Demo struct{}

// Preload loads the given files into memory before the application begins.
func (*Demo) Preload() {
	engo.Files.Load("player.png")
}

// Setup will setup the scene including all the of the systems and components.
func (*Demo) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	// Open the aseprite JSON file and play the rightaction by default.
	ase, _ := aseprite.Open("assets/player.json")
	ase.Play("rightaction")

	// Create a new empty player.
	player := &Player{
		ase:         ase,
		BasicEntity: ecs.NewBasic(),
	}
	// Create a spritesheet using the file and its frame boundaries.
	player.spritesheet = common.NewSpritesheetFromFile(
		"player.png",
		int(player.ase.FrameBoundaries().Width),
		int(player.ase.FrameBoundaries().Height),
	)
	// Create a renderable object using the common.Spritesheet.
	player.RenderComponent = common.RenderComponent{
		Drawable: player.spritesheet.Drawable(0),
		Scale:    engo.Point{X: entityScale, Y: entityScale},
	}
	// Create a space component place the sprite in the scene.
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    player.spritesheet.Width() * player.RenderComponent.Scale.X,
		Height:   player.spritesheet.Height() * player.RenderComponent.Scale.Y,
	}

	// Add systems.
	w.AddSystem(player)
	w.AddSystem(&common.RenderSystem{})

	// Loop through the added systems to add them.
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}
}

// Type returns the name of the scene.
func (*Demo) Type() string {
	return "Demo"
}

// Player is used to display the sprite itself. It also contains the
// aseprite file and the other components needed to display with engo.
type Player struct {
	spritesheet *common.Spritesheet
	ase         *aseprite.File

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Update gets called every frame and updates the Drawable object to whatever
// frame the aseprite file is on.
func (p *Player) Update(dt float32) {
	p.ase.Update(dt)
	p.Drawable = p.spritesheet.Drawable(p.ase.CurrentFrame)
}

// Remove will remove the Drawable object from the scene.
func (p *Player) Remove(ecs.BasicEntity) {
	p.Drawable.Close()
}
