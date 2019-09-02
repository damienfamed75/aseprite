package aseprite

type MetaData struct {
	App     string      `json:"app"`
	Version string      `json:"version"`
	Image   string      `json:"image"`
	Format  string      `json:"format"`
	Size    WidthHeight `json:"size"`
	Scale   string      `json:"scale"`
	Tags    []FrameTag  `json:"frameTags,omitempty"`
	Layers  []Layer     `json:"layers,omitempty"`
	Slices  []Slice     `json:"slices,omitempty"`
}

type FrameTag struct {
	Name string `json:"name"`
	From int    `json:"from"`
	To   int    `json:"to"`
	// TODO Use stringer for direction modes.
	Direction string `json:"direction"`
}

type Layer struct {
	Name    string `json:"name"`
	Opacity int    `json:"opacity"`
	// TODO Use stringer for blend modes.
	BlendMode string `json:"blendMode"`
}

type Slice struct {
	Name string `json:"name"`
	// TODO Color to image.Color
	Color string     `json:"color"`
	Keys  []SliceKey `json:"keys"`
}

type SliceKey struct {
	FrameNum int      `json:"frame"`
	Bounds   Boundary `json:"bounds"`
	Center   Boundary `json:"center"`
	Pivot    Point    `json:"pivot"`
}
