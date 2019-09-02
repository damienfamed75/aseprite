package aseprite

// File is the final parsed aseprite file.
// This includes helpful methods to interact and get information about
// the animations and its frames.
type File struct {
	Frames FrameData `json:"frames"`
	Meta   MetaData  `json:"meta"`
}
