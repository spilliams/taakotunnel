package model

const (
	TunnelNotConnected int = iota
	TunnelTopLeft
	TunnelTopRight
	TunnelRightTop
	TunnelRightBottom
	TunnelBottomRight
	TunnelBottomLeft
	TunnelLeftBottom
	TunnelLeftTop
)

const (
	DirectionNorth int = iota
	DirectionWest
	DirectionSouth
	DirectionEast
)
