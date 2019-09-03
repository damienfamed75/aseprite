package aseprite

import (
	"encoding/json"

	"gitlab.com/c0b/go-ordered-json"
)

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
	o := ordered.NewOrderedMap()
	// Unmarshal the JSON into an ordered map.
	// If there wasn't an error while unmarshalling then process the map.
	if err := json.Unmarshal(data, &o); err == nil {
		d.frameMap = make(map[string]Frame)
		d.IsMap = true

		// Iterate through the map in incremental order.
		iter := o.EntriesIter()
		for {
			// Grab the next pair of items.
			pair, ok := iter()
			if !ok { // If there wasn't another pair then break.
				break
			}

			// Create a new Frame object to append.
			var f Frame
			// Marshal the interface{} into a JSON.
			bytes, err := json.Marshal(pair.Value)
			if err != nil {
				return err
			}
			// Unmarshal the JSON into the Frame.
			if err := json.Unmarshal(bytes, &f); err != nil {
				return err
			}
			// Set the FileName of the Frame.
			f.FileName = pair.Key
			// Fix the duration to work correctly when playing back.
			f.Duration = f.Duration / 1000
			// Append the frame to the slice and map of frames.
			d.frameSlice = append(d.frameSlice, f)
			d.frameMap[pair.Key] = f
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
