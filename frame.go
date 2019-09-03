package aseprite

import "encoding/json"

// FrameData contains all the frames in the animation and their data.
type FrameData struct {
	frameSlice []Frame          // used for reference.
	frameMap   map[string]Frame // used primarily.
	IsMap      bool             `json:"-"`
}

// Frame contains the data about a frame.
type Frame struct {
	FileName         string      `json:"filename,omitempty"`
	FrameBoundaries  Boundary    `json:"frame"`
	Rotated          bool        `json:"rotated"`
	Trimmed          bool        `json:"trimmmed"`
	SpriteSourceSize Boundary    `json:"spriteSourceSize"`
	SourceSize       WidthHeight `json:"sourceSize"`
	Duration         float32     `json:"duration"`
}

// LenFrames return the length of the number of frames in the aseprite file.
func (d *FrameData) LenFrames() int {
	return len(d.frameSlice)
}

// FrameAtIndex uses the local copy of the sliced frames and searches for the frame
// on the nth index.
func (d *FrameData) FrameAtIndex(index int) Frame {
	return d.frameSlice[index]
}

// FrameAtKey uses the local copy of the mapped frames and searches for the frame
// assigned to the given key.
func (d *FrameData) FrameAtKey(key string) Frame {
	return d.frameMap[key]
}

// FrameSlice returns a reference to the sliced version of the frames.
func (d *FrameData) FrameSlice() []Frame {
	return d.frameSlice
}

// FrameMap returns a reference to the mapped version of the frames.
func (d *FrameData) FrameMap() map[string]Frame {
	return d.frameMap
}

// UnmarshalJSON handles the problem of different layouts in the JSON file.
// Sometimes the file can be a map for the frames and sometimes it can be a slice.
func (d *FrameData) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &d.frameMap); err == nil {
		d.IsMap = true
		for k, f := range d.frameMap {
			f.FileName = k
			f.Duration = f.Duration / 1000
			d.frameSlice = append(d.frameSlice, f)
		}
		return nil
	}

	if err := json.Unmarshal(data, &d.frameSlice); err != nil {
		return err
	}

	d.frameMap = make(map[string]Frame)
	for _, f := range d.frameSlice {
		f.Duration = f.Duration / 1000
		d.frameMap[f.FileName] = f
	}

	return nil
}
