package render

import (
	"fmt"
	"strings"

	"github.com/enotinc/cubeit/internal/cube"
)

const (
	color_white  = "\033[37m"
	color_yellow = "\033[33m"
	color_green  = "\033[32m"
	color_blue   = "\033[34m"
	color_orange = "\033[38;5;208m"
	color_red    = "\033[31m"
	color_reset  = "\033[0m"
)

const (
	tile  = "██"
	empty = "      "
	//......|123456|
)

func Render(c *cube.Cube) {
	var r strings.Builder

	scramble := fmt.Sprintf(" %s\n\n", c.Scramble)
	r.WriteString(scramble)

	for i := range 3 {
		for row := range 3 {
			switch i {
			case 0:
				up := [3]cube.Color{c.Faces[cube.UP][row*3], c.Faces[cube.UP][row*3+1], c.Faces[cube.UP][row*3+2]}

				line := fmt.Sprintf("%s  %s\n", empty, fl(up))
				r.WriteString(line)
			case 1:
				left := [3]cube.Color{c.Faces[cube.LEFT][row*3], c.Faces[cube.LEFT][row*3+1], c.Faces[cube.LEFT][row*3+2]}
				front := [3]cube.Color{c.Faces[cube.FRONT][row*3], c.Faces[cube.FRONT][row*3+1], c.Faces[cube.FRONT][row*3+2]}
				right := [3]cube.Color{c.Faces[cube.RIGHT][row*3], c.Faces[cube.RIGHT][row*3+1], c.Faces[cube.RIGHT][row*3+2]}
				back := [3]cube.Color{c.Faces[cube.BACK][row*3], c.Faces[cube.BACK][row*3+1], c.Faces[cube.BACK][row*3+2]}

				line := fmt.Sprintf("%s  %s  %s  %s\n", fl(left), fl(front), fl(right), fl(back))
				r.WriteString(line)

			case 2:
				down := [3]cube.Color{c.Faces[cube.DOWN][row*3], c.Faces[cube.DOWN][row*3+1], c.Faces[cube.DOWN][row*3+2]}

				line := fmt.Sprintf("%s  %s\n", empty, fl(down))
				r.WriteString(line)
			}
		}
		r.WriteString("\n")
	}

	fmt.Print(r.String())
}

var colors = map[cube.Color]string{
	cube.White:  color_white,
	cube.Yellow: color_yellow,
	cube.Green:  color_green,
	cube.Blue:   color_blue,
	cube.Orange: color_orange,
	cube.Red:    color_red,
}

// fl - face line
func fl(fl [3]cube.Color) string {
	var face strings.Builder

	face.WriteString(fmt.Sprintf("%s%s%s%s%s%s%s", colors[fl[0]], tile, colors[fl[1]], tile, colors[fl[2]], tile, color_reset))

	return face.String()
}
