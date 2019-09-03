package aseprite

import "encoding/json"

type FrameData struct {
	frameSlice []Frame
	frameMap   map[string]Frame
	IsMap      bool `json:"-"`
}

func (d *FrameData) Frames() map[string]Frame {
	return d.frameMap
}

func (d *FrameData) FrameAt(index int) Frame {
	return d.frameSlice[index]
}

type Frame struct {
	FileName         string      `json:"filename,omitempty"`
	FrameBoundaries  Boundary    `json:"frame"`
	Rotated          bool        `json:"rotated"`
	Trimmed          bool        `json:"trimmmed"`
	SpriteSourceSize Boundary    `json:"spriteSourceSize"`
	SourceSize       WidthHeight `json:"sourceSize"`
	Duration         int         `json:"duration"`
}

func (d *FrameData) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &d.frameMap); err == nil {
		d.IsMap = true
		for k, f := range d.frameMap {
			f.FileName = k
			d.frameSlice = append(d.frameSlice, f)
		}
		return nil
	}

	if err := json.Unmarshal(data, &d.frameSlice); err != nil {
		return err
	}

	d.frameMap = make(map[string]Frame)
	for _, f := range d.frameSlice {
		d.frameMap[f.FileName] = f
	}

	return nil
}
