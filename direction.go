package aseprite

//go:generate stringer -type=Direction -linecomment

// Direction represents the playback directions for an animation.
type Direction uint8

const (
	// PlayForward playsback frames in incremental order.
	// Frames [0, 1, 2, 3, 4, 5]
	// Order  [0, 1, 2, 3, 4, 5]
	PlayForward Direction = iota // forward
	// PlayReverse playsback frames in decremental order.
	// Frames [0, 1, 2, 3, 4, 5]
	// Order  [5, 4, 3, 2, 1, 0]
	PlayReverse // reverse
	// PlayPingPong plays frames in incremental order, then once reaching the
	// last frame the animation will play in decremental order.
	// Frames [0, 1, 2, 3, 4, 5]
	// Order  [0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0, ...]
	PlayPingPong // pingpong
)

// UnmarshalJSON unmarshals a byte slice representing the string values
// of the data and then sets the value to a Direction.
func (d *Direction) UnmarshalJSON(data []byte) error {
	*d = DirectionFromString(string(data))

	return nil
}

// DirectionFromString returns a Direction based on the value of the string.
// Note: This function assumes that the string is lowercased correctly.
func DirectionFromString(str string) Direction {
	switch str {
	case PlayReverse.String():
		return PlayReverse
	case PlayPingPong.String():
		return PlayPingPong
	case PlayForward.String():
		fallthrough
	default:
		return PlayForward
	}
}
