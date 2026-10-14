package cube

type color string

const (
	Yellow = "y"
	White  = "w"
	Green  = "g"
	Orange = "o"
	Blue   = "b"
	Red    = "r"
)

type dir int

const (
	_ dir = iota - 1
	UP
	DOWN
	LEFT
	RIGHT
	FRONT
	BACK
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
