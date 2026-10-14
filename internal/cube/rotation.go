package cube

import "fmt"

func (c *Cube) rotate(m Move) error {
	switch m {
	case U:
		old := c.faces
		c.faces[up] = matrixClockwise(old[up])
		copy(c.faces[front][0:3], old[right][0:3])
		copy(c.faces[right][0:3], old[back][0:3])
		copy(c.faces[back][0:3], old[left][0:3])
		copy(c.faces[left][0:3], old[front][0:3])

	case Up:
		old := c.faces
		c.faces[up] = matrixCounterClockwise(old[up])
		copy(c.faces[front][0:3], old[left][0:3])
		copy(c.faces[left][0:3], old[back][0:3])
		copy(c.faces[back][0:3], old[right][0:3])
		copy(c.faces[right][0:3], old[front][0:3])

	case D:
		old := c.faces
		c.faces[down] = matrixClockwise(old[down])
		copy(c.faces[front][6:9], old[left][6:9])
		copy(c.faces[left][6:9], old[back][6:9])
		copy(c.faces[back][6:9], old[right][6:9])
		copy(c.faces[right][6:9], old[front][6:9])

	case Dp:
		old := c.faces
		c.faces[down] = matrixCounterClockwise(old[down])
		copy(c.faces[front][6:9], old[right][6:9])
		copy(c.faces[right][6:9], old[back][6:9])
		copy(c.faces[back][6:9], old[left][6:9])
		copy(c.faces[left][6:9], old[front][6:9])

	default:
		return fmt.Errorf("Unknown move %+v", m)
	}

	return nil
}

// lazy clockwise matrix rotation
func matrixClockwise(old face) face {
	f := face{
		old[6], old[3], old[0],
		old[7], old[4], old[1],
		old[8], old[5], old[2],
	}

	return f
}

// lazy counterclockwise matrix rotation
func matrixCounterClockwise(old face) face {
	f := face{
		old[2], old[5], old[8],
		old[1], old[4], old[7],
		old[0], old[3], old[6],
	}

	return f
}
