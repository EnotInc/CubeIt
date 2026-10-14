package cube

type face [9]Color

type Cube struct {
	Scramble string
	Faces    [6]face
}

func New() *Cube {
	c := &Cube{}

	c.Faces[int(UP)] = [9]Color{White, White, White, White, White, White, White, White, White}
	c.Faces[int(DOWN)] = [9]Color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow}
	c.Faces[int(FRONT)] = [9]Color{Green, Green, Green, Green, Green, Green, Green, Green, Green}
	c.Faces[int(BACK)] = [9]Color{Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue}
	c.Faces[int(LEFT)] = [9]Color{Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange}
	c.Faces[int(RIGHT)] = [9]Color{Red, Red, Red, Red, Red, Red, Red, Red, Red}

	return c
}
