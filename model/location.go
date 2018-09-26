package model

type Location struct {
	row    int // 0-2
	col    int // 0-2
	tunnel int // 0-8
	end    bool
}
