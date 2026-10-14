package cube

type face [9]color

type Cube struct {
	Faces [6]face
}

func New() *Cube {
	c := &Cube{}

	c.Faces[int(UP)] = [9]color{White, White, White, White, White, White, White, White, White}
	c.Faces[int(DOWN)] = [9]color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow}
	c.Faces[int(FRONT)] = [9]color{Green, Green, Green, Green, Green, Green, Green, Green, Green}
	c.Faces[int(BACK)] = [9]color{Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue}
	c.Faces[int(LEFT)] = [9]color{Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange}
	c.Faces[int(RIGHT)] = [9]color{Red, Red, Red, Red, Red, Red, Red, Red, Red}

	return c
}

func (c *Cube) Scamble() {
	panic("not implemented yet")
}

func (c *Cube) ScambleFrom(s string) {
	panic("not implemented yet")
}
