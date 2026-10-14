package cube

type color string

const (
	yellow = "y"
	white  = "w"
	green  = "g"
	orange = "o"
	blue   = "b"
	red    = "r"
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

type Move string

const (
	// Up, Down, Front, Back, Right, Left
	// p - stands for prime (counterclockwise)
	U  = "U"
	Up = "U'"
	D  = "D"
	Dp = "D'"
	F  = "F"
	Fp = "F'"
	Ba = "B"
	Bp = "B'"
	R  = "R"
	Rp = "R'"
	L  = "L"
	Lp = "L'"
)
