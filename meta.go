package aseprite

import "image"

// MetaData is the data about the overall animation including
// animation playback directions, layer opacity, and pivot points.
type MetaData struct {
	App        string      `json:"app"`
	Version    string      `json:"version"`
	Image      string      `json:"image"`
	Format     string      `json:"format"`
	Size       WidthHeight `json:"size"`
	Scale      string      `json:"scale"`
	Animations []Animation `json:"frameTags,omitempty"`
	Layers     []Layer     `json:"layers,omitempty"`
	Slices     []Slice     `json:"slices,omitempty"`
}

// LenAnimations returns the length of the Animations slice.
func (m *MetaData) LenAnimations() int {
	return len(m.Animations)
}

// LenLayers returns the length of the Layers slice.
func (m *MetaData) LenLayers() int {
	return len(m.Layers)
}

// LenSlices returns the length of the Slices slice.
func (m *MetaData) LenSlices() int {
	return len(m.Slices)
}

// Animation holds the metadata about the playing animation
// including its beginning frame (From) and ending frame (To).
type Animation struct {
	// Name of the animation.
	Name string `json:"name"`
	// The beginning frame of the animation.
	From int `json:"from"`
	// The ending frame of the animation.
	To int `json:"to"`
	// Animation's playback direction.
	// Example:
	//  - PlayForward
	//  - PlayReverse
	//  - PlayPingPong
	Direction Direction `json:"direction"`
}

// Layer contains the data about the individual animation layer
// including its name, opacity, and blending mode.
// Layers can be used to stack on top of each other to create
// unique affects on the animation itself.
type Layer struct {
	Name      string `json:"name"`
	Opacity   int    `json:"opacity"`
	BlendMode string `json:"blendMode"`
}

// Slice is used for slicing up an image in a layer or animation in aseprite.
// In this struct it contains the information about the slice itself including
// its keys, color, and the name of the slice.
type Slice struct {
	Name  string     `json:"name"`
	Color string     `json:"color"`
	Keys  []SliceKey `json:"keys"`
}

// SliceKey stores the information about each slice including its
// frame number, the boundaries of the slice, and more.
type SliceKey struct {
	FrameNum int         `json:"frame"`
	Bounds   Boundary    `json:"bounds"`
	Center   Boundary    `json:"center"`
	Pivot    image.Point `json:"pivot"`
}
