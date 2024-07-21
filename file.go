package aseprite

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// File is the final parsed aseprite file.
// This includes helpful methods to interact and get information about
// the animations and its frames.
type File struct {
	Frames FrameData `json:"frames"`
	Meta   MetaData  `json:"meta"`
	AnimationInfo
}

// Open will remove some of the pain of reading the data from the aseprite
// JSON file and loading it into memory then giving it over here to turn into
// a file. Instead this will open the filepath and then read the bytes and return
// an aseprite File.
func Open(filePath string) (*File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return NewFile(bytes)
}

// NewFile takes the file's data and unmarshals it into a new File.
// What's the point in using this? Maybe saving an extra import?
// I'm not really sure at this point.
func NewFile(data []byte) (*File, error) {
	f := &File{AnimationInfo: setupAnimationInfo()}

	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}

	return f, nil
}

// Play will mark an animation in the list of animations listed in the aseprite
// file to be used for the Update method.
func (f *File) Play(animation string) error {
	anim := f.Animation(animation)
	if anim == nil {
		return errorAnimationNotFound.withParams(animation)
	}

	if f.CurrentAnimation == nil || *anim != *f.CurrentAnimation {
		f.playAnimation(anim)
	}

	return nil
}

// FrameBoundaries will return the current frame's bounding box.
func (f *File) FrameBoundaries() Boundary {
	if f.CurrentAnimation != nil {
		return f.Frames.FrameAtIndex(f.CurrentFrame).FrameBoundaries
	}

	return Boundary{}
}

// Animation will search the slices in the aseprite file
// and return any animation that matches the name provided.
// Note: Animation is a duplicate of GetSlice is almost every way.
// This is because of the limitations of Golang that I cannot address.
// There are no generics in the language yet that would allow me to
// prevent this code duplication.
func (f *File) Animation(animation string) *Animation {
	for _, anim := range f.Meta.Animations {
		if anim.Name == animation {
			return &anim
		}
	}

	return nil
}

// Slice will search the slices in the aseprite file
// and return any slice that matches the name provided.
// Note: Slice is a duplicate of GetAnimation is almost every way.
// This is because of the limitations of Golang that I cannot address.
// There are no generics in the language yet that would allow me to
// prevent this code duplication.
func (f *File) Slice(slice string) *Slice {
	for _, s := range f.Meta.Slices {
		if s.Name == slice {
			return &s
		}
	}

	return nil
}

// HasAnimation returns a boolean value that represents whether the provided
// animation is contained in the aseprite file.
func (f *File) HasAnimation(animation string) bool {
	return f.Animation(animation) != nil
}

// HasSlice returns a boolean value that represents whether the provided
// slice is contained in the aseprite file.
func (f *File) HasSlice(slice string) bool {
	return f.Slice(slice) != nil
}

// Update is used for a Game Loop in such an OpenGL style workflow.
// This can be helpful for tracking the frame of the animation.
func (f *File) Update(dt float32) {
	f.PrevFrame = f.PrevCurrentFrame
	f.PrevCurrentFrame = f.CurrentFrame
	f.animationFinished = false

	if f.CurrentAnimation != nil {
		// Increment the frame counter based on delta time.
		// Note: Truncate multiplication.
		f.FrameCounter += dt * f.PlaySpeed

		// If the frame counter is greater than the expected frame duration
		// then increment or decrement the current frame being displayed.
		if f.FrameCounter > float32(f.Frames.FrameAtIndex(f.CurrentFrame).Duration) {
			f.advanceFrame()
		}

		// update the values to make sure that the animations are being
		// finished or ping ponged back and forth.
		f.updateAnimationValues()
	}
}
