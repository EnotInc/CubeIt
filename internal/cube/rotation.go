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

	case F:
		old := c.Faces
		c.Faces[FRONT] = matrixClockwise(old[FRONT])
		c.Faces[UP][6] = old[LEFT][8]
		c.Faces[UP][7] = old[LEFT][5]
		c.Faces[UP][8] = old[LEFT][2]

		c.Faces[LEFT][8] = old[DOWN][0]
		c.Faces[LEFT][5] = old[DOWN][1]
		c.Faces[LEFT][2] = old[DOWN][2]

		c.Faces[DOWN][0] = old[RIGHT][0]
		c.Faces[DOWN][1] = old[RIGHT][3]
		c.Faces[DOWN][2] = old[RIGHT][6]

		c.Faces[RIGHT][0] = old[UP][6]
		c.Faces[RIGHT][3] = old[UP][7]
		c.Faces[RIGHT][6] = old[UP][8]

	case Fp:
		old := c.Faces
		c.Faces[FRONT] = matrixCounterClockwise(old[FRONT])
		c.Faces[UP][6] = old[RIGHT][0]
		c.Faces[UP][7] = old[RIGHT][3]
		c.Faces[UP][8] = old[RIGHT][6]

		c.Faces[LEFT][8] = old[UP][6]
		c.Faces[LEFT][5] = old[UP][7]
		c.Faces[LEFT][2] = old[UP][8]

		c.Faces[DOWN][0] = old[LEFT][2]
		c.Faces[DOWN][1] = old[LEFT][5]
		c.Faces[DOWN][2] = old[LEFT][8]

		c.Faces[RIGHT][0] = old[DOWN][0]
		c.Faces[RIGHT][3] = old[DOWN][1]
		c.Faces[RIGHT][6] = old[DOWN][2]

	case B:
		old := c.Faces
		c.Faces[BACK] = matrixClockwise(old[BACK])

		c.Faces[UP][0] = old[RIGHT][2]
		c.Faces[UP][1] = old[RIGHT][5]
		c.Faces[UP][2] = old[RIGHT][8]

		c.Faces[RIGHT][2] = old[DOWN][6]
		c.Faces[RIGHT][5] = old[DOWN][7]
		c.Faces[RIGHT][8] = old[DOWN][8]

		c.Faces[DOWN][6] = old[LEFT][0]
		c.Faces[DOWN][7] = old[LEFT][3]
		c.Faces[DOWN][8] = old[LEFT][6]

		c.Faces[LEFT][0] = old[UP][0]
		c.Faces[LEFT][3] = old[UP][1]
		c.Faces[LEFT][6] = old[UP][2]

	case Bp:
		old := c.Faces
		c.Faces[BACK] = matrixCounterClockwise(old[BACK])

		c.Faces[UP][0] = old[LEFT][0]
		c.Faces[UP][1] = old[LEFT][3]
		c.Faces[UP][2] = old[LEFT][6]

		c.Faces[LEFT][0] = old[DOWN][6]
		c.Faces[LEFT][3] = old[DOWN][7]
		c.Faces[LEFT][6] = old[DOWN][8]

		c.Faces[DOWN][6] = old[RIGHT][2]
		c.Faces[DOWN][7] = old[RIGHT][5]
		c.Faces[DOWN][8] = old[RIGHT][8]

		c.Faces[RIGHT][2] = old[UP][0]
		c.Faces[RIGHT][5] = old[UP][1]
		c.Faces[RIGHT][8] = old[UP][2]

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
