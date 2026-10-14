package internal

type color int

const (
	_ color = iota
	yellow
	white
	green
	orange
	blue
	red
)

type dir int

const (
	_ dir = iota - 1
	up
	down
	left
	right
	front
	back
)

type Move int

const (
	_ Move = iota
	// Up, Down, Front, Back, Right, Left
	// p - stands for prime (counterclockwise)

	U
	Up
	D
	Dp
	F
	Fp
	B
	Bp
	R
	Rp
	L
	Lp
)
