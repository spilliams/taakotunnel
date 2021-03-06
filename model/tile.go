package model

import "fmt"

type Tile struct {
	letter   string
	tunnels  map[int]int
	rotation int
}

// NewTile creates a new tile, with the given letter and with the connections
// of the given tunnels. The new tile's rotation is Top in the North (0)
func NewTile(tunnels []Tunnel, letter string) (*Tile, error) {
	tunnelMap := make(map[int]int, 9)

	t := Tile{tunnels: tunnelMap, rotation: DirectionNorth, letter: letter}

	for _, tunnel := range tunnels {
		e := t.connect(tunnel.In, tunnel.Out)
		if e != nil {
			return nil, e
		}
	}

	return &t, nil
}

// connect will connect two tunnels on a tile to each other.
func (t *Tile) connect(inlet, outlet int) error {
	if inlet <= TunnelNotConnected || inlet > TunnelLeftTop {
		return fmt.Errorf("inlet out of bounds: %v", inlet)
	}
	if outlet <= TunnelNotConnected || outlet > TunnelLeftTop {
		return fmt.Errorf("outlet out of bounds: %v", outlet)
	}

	t.tunnels[inlet] = outlet
	t.tunnels[outlet] = inlet

	return nil
}

// RotateTopTo rotates the tile such that the top is facing the given direction
// 0: Top in the North
// 1: Top in the West
// 2: Top in the South
// 3: Top in the East
func (t *Tile) RotateTopTo(rotation int) error {
	if rotation < DirectionNorth || rotation > DirectionEast {
		return fmt.Errorf("rotation out of bounds %v", rotation)
	}

	t.rotation = rotation
	return nil
}

// Follow will return the outlet of the tunnel given by the `inlet`,
// while respecting tile rotation.
//
// For example: Following an inlet of `TunnelTopLeft` on a tile rotated to
// `DirectionWest` will "enter" the tile at the absolute top left,
// which corresponds to the tile's internal index for `TunnelLeftBottom`.
func (t *Tile) Follow(inlet int) (int, error) {
	if inlet <= TunnelNotConnected || inlet > TunnelLeftTop {
		return TunnelNotConnected, fmt.Errorf("inlet out of bounds %v", inlet)
	}

	// inlet = inlet + t.rotation*2
	// if inlet > 8 {
	// 	inlet = inlet - 8
	// }

	return t.tunnels[inlet], nil
}

// NextTunnel translates an outlet tunnel code to an inlet tunnel code
func NextTunnel(inlet int) (int, error) {
	if inlet <= TunnelNotConnected || inlet > TunnelLeftTop {
		return TunnelNotConnected, fmt.Errorf("inlet out of bounds %v", inlet)
	}
	/*
	    6 5
	    1 2
	 3 8   3 8
	 4 7   4 7
	    6 5
	    1 2
	*/
	io := map[int]int{
		1: 6,
		2: 5,
		3: 8,
		4: 7,
		5: 2,
		6: 1,
		7: 4,
		8: 3,
	}
	return io[inlet], nil
}

func (t *Tile) Letter() string {
	return t.letter
}

func (t *Tile) Rotation() int {
	return t.rotation
}
