package cube

import "fmt"

func (c *Cube) rotate(m Move) error {
	switch m {
	case U:
		old := c.Faces
		c.Faces[UP] = matrixClockwise(old[UP])
		copy(c.Faces[FRONT][0:3], old[RIGHT][0:3])
		copy(c.Faces[RIGHT][0:3], old[BACK][0:3])
		copy(c.Faces[BACK][0:3], old[LEFT][0:3])
		copy(c.Faces[LEFT][0:3], old[FRONT][0:3])

	case Up:
		old := c.Faces
		c.Faces[UP] = matrixCounterClockwise(old[UP])
		copy(c.Faces[FRONT][0:3], old[LEFT][0:3])
		copy(c.Faces[LEFT][0:3], old[BACK][0:3])
		copy(c.Faces[BACK][0:3], old[RIGHT][0:3])
		copy(c.Faces[RIGHT][0:3], old[FRONT][0:3])

	case D:
		old := c.Faces
		c.Faces[DOWN] = matrixClockwise(old[DOWN])
		copy(c.Faces[FRONT][6:9], old[LEFT][6:9])
		copy(c.Faces[LEFT][6:9], old[BACK][6:9])
		copy(c.Faces[BACK][6:9], old[RIGHT][6:9])
		copy(c.Faces[RIGHT][6:9], old[FRONT][6:9])

	case Dp:
		old := c.Faces
		c.Faces[DOWN] = matrixCounterClockwise(old[DOWN])
		copy(c.Faces[FRONT][6:9], old[RIGHT][6:9])
		copy(c.Faces[RIGHT][6:9], old[BACK][6:9])
		copy(c.Faces[BACK][6:9], old[LEFT][6:9])
		copy(c.Faces[LEFT][6:9], old[FRONT][6:9])

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
