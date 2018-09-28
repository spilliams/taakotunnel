package model

type Location struct {
	Row    int // 0-2
	Col    int // 0-2
	Tunnel int // 0-8
	End    bool
}
