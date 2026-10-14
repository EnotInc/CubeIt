package cube

import "fmt"

// NOTE: I could fk up some indexes...
func (c *Cube) rotate(m Move) error {
	old := c.Faces
	switch m {
	case U:
		c.Faces[UP] = matrixClockwise(old[UP])
		copy(c.Faces[FRONT][0:3], old[RIGHT][0:3])
		copy(c.Faces[RIGHT][0:3], old[BACK][0:3])
		copy(c.Faces[BACK][0:3], old[LEFT][0:3])
		copy(c.Faces[LEFT][0:3], old[FRONT][0:3])

	case Up:
		c.Faces[UP] = matrixCounterClockwise(old[UP])
		copy(c.Faces[FRONT][0:3], old[LEFT][0:3])
		copy(c.Faces[LEFT][0:3], old[BACK][0:3])
		copy(c.Faces[BACK][0:3], old[RIGHT][0:3])
		copy(c.Faces[RIGHT][0:3], old[FRONT][0:3])

	case D:
		c.Faces[DOWN] = matrixClockwise(old[DOWN])
		copy(c.Faces[FRONT][6:9], old[LEFT][6:9])
		copy(c.Faces[LEFT][6:9], old[BACK][6:9])
		copy(c.Faces[BACK][6:9], old[RIGHT][6:9])
		copy(c.Faces[RIGHT][6:9], old[FRONT][6:9])

	case Dp:
		c.Faces[DOWN] = matrixCounterClockwise(old[DOWN])
		copy(c.Faces[FRONT][6:9], old[RIGHT][6:9])
		copy(c.Faces[RIGHT][6:9], old[BACK][6:9])
		copy(c.Faces[BACK][6:9], old[LEFT][6:9])
		copy(c.Faces[LEFT][6:9], old[FRONT][6:9])

	case F:
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

	case R:
		c.Faces[RIGHT] = matrixClockwise(old[RIGHT])

		c.Faces[UP][2] = old[FRONT][2]
		c.Faces[UP][5] = old[FRONT][5]
		c.Faces[UP][8] = old[FRONT][8]

		c.Faces[FRONT][2] = old[DOWN][2]
		c.Faces[FRONT][5] = old[DOWN][5]
		c.Faces[FRONT][8] = old[DOWN][8]

		c.Faces[DOWN][2] = old[BACK][0]
		c.Faces[DOWN][5] = old[BACK][3]
		c.Faces[DOWN][8] = old[BACK][6]

		c.Faces[BACK][0] = old[UP][2]
		c.Faces[BACK][3] = old[UP][5]
		c.Faces[BACK][6] = old[UP][8]

	case Rp:
		c.Faces[RIGHT] = matrixCounterClockwise(old[RIGHT])

		c.Faces[UP][2] = old[BACK][0]
		c.Faces[UP][5] = old[BACK][3]
		c.Faces[UP][8] = old[BACK][6]

		c.Faces[BACK][0] = old[DOWN][2]
		c.Faces[BACK][3] = old[DOWN][5]
		c.Faces[BACK][6] = old[DOWN][8]

		c.Faces[DOWN][2] = old[FRONT][2]
		c.Faces[DOWN][5] = old[FRONT][5]
		c.Faces[DOWN][8] = old[FRONT][8]

		c.Faces[FRONT][2] = old[UP][2]
		c.Faces[FRONT][5] = old[UP][5]
		c.Faces[FRONT][8] = old[UP][8]

	case L:
		c.Faces[LEFT] = matrixClockwise(old[LEFT])

		c.Faces[UP][0] = old[BACK][2]
		c.Faces[UP][3] = old[BACK][5]
		c.Faces[UP][6] = old[BACK][8]

		c.Faces[BACK][2] = old[DOWN][0]
		c.Faces[BACK][5] = old[DOWN][3]
		c.Faces[BACK][8] = old[DOWN][6]

		c.Faces[DOWN][0] = old[FRONT][0]
		c.Faces[DOWN][3] = old[FRONT][3]
		c.Faces[DOWN][6] = old[FRONT][6]

		c.Faces[FRONT][0] = old[UP][0]
		c.Faces[FRONT][3] = old[UP][3]
		c.Faces[FRONT][6] = old[UP][6]

	case Lp:
		c.Faces[LEFT] = matrixCounterClockwise(old[LEFT])

		c.Faces[UP][0] = old[FRONT][0]
		c.Faces[UP][3] = old[FRONT][3]
		c.Faces[UP][6] = old[FRONT][6]

		c.Faces[FRONT][0] = old[DOWN][0]
		c.Faces[FRONT][3] = old[DOWN][3]
		c.Faces[FRONT][6] = old[DOWN][6]

		c.Faces[DOWN][0] = old[BACK][2]
		c.Faces[DOWN][3] = old[BACK][5]
		c.Faces[DOWN][6] = old[BACK][8]

		c.Faces[BACK][2] = old[UP][0]
		c.Faces[BACK][5] = old[UP][3]
		c.Faces[BACK][8] = old[UP][6]

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
