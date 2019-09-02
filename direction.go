package aseprite

type Direction uint8

const (
	Forward Direction = iota
	Reverse
	PingPong
)
