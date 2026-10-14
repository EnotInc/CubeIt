package cube

type face [9]color

type Cube struct {
	faces [6]face
}

func New() *Cube {
	c := &Cube{}

	c.faces[int(up)] = [9]color{white, white, white, white, white, white, white, white, white}
	c.faces[int(down)] = [9]color{yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow}
	c.faces[int(front)] = [9]color{green, green, green, green, green, green, green, green, green}
	c.faces[int(back)] = [9]color{blue, blue, blue, blue, blue, blue, blue, blue, blue}
	c.faces[int(left)] = [9]color{orange, orange, orange, orange, orange, orange, orange, orange, orange}
	c.faces[int(right)] = [9]color{red, red, red, red, red, red, red, red, red}

	return c
}

func (c *Cube) Scamble() {
	panic("not implemented yet")
}

func (c *Cube) ScambleFrom(s string) {
	panic("not implemented yet")
}
