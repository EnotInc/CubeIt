package cube

import (
	"testing"
)

var allYellow = [9]color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow}
var allWhite = [9]color{White, White, White, White, White, White, White, White, White}
var allGreen = [9]color{Green, Green, Green, Green, Green, Green, Green, Green, Green}
var allOrange = [9]color{Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange}
var allBlue = [9]color{Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue}
var allRed = [9]color{Red, Red, Red, Red, Red, Red, Red, Red, Red}

func testCube(t *testing.T) *Cube {
	c := New()
	testDefaultCube(t, c)
	return c
}

func testDefaultCube(t *testing.T, c *Cube) {
	testFace(t, c.Faces[UP], allWhite)
	testFace(t, c.Faces[DOWN], allYellow)
	testFace(t, c.Faces[FRONT], allGreen)
	testFace(t, c.Faces[BACK], allBlue)
	testFace(t, c.Faces[LEFT], allOrange)
	testFace(t, c.Faces[RIGHT], allRed)
}

func TestCubeUMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(U)
	if err != nil {
		t.Errorf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[UP], [9]color{White, White, White, White, White, White, White, White, White})
	testFace(t, c.Faces[FRONT], [9]color{Red, Red, Red, Green, Green, Green, Green, Green, Green})
	testFace(t, c.Faces[RIGHT], [9]color{Blue, Blue, Blue, Red, Red, Red, Red, Red, Red})
	testFace(t, c.Faces[BACK], [9]color{Orange, Orange, Orange, Blue, Blue, Blue, Blue, Blue, Blue})
	testFace(t, c.Faces[LEFT], [9]color{Green, Green, Green, Orange, Orange, Orange, Orange, Orange, Orange})

	err = c.rotate(Up)
	if err != nil {
		t.Errorf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestCubeDMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(D)
	if err != nil {
		t.Errorf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[DOWN], [9]color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow})
	testFace(t, c.Faces[FRONT], [9]color{Green, Green, Green, Green, Green, Green, Orange, Orange, Orange})
	testFace(t, c.Faces[RIGHT], [9]color{Red, Red, Red, Red, Red, Red, Green, Green, Green})
	testFace(t, c.Faces[BACK], [9]color{Blue, Blue, Blue, Blue, Blue, Blue, Red, Red, Red})
	testFace(t, c.Faces[LEFT], [9]color{Orange, Orange, Orange, Orange, Orange, Orange, Blue, Blue, Blue})

	err = c.rotate(Dp)
	if err != nil {
		t.Errorf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func testFace(t *testing.T, f face, c [9]color) {
	if f != c {
		t.Errorf("wrong face rotation.\nWant\t%+v\nGot\t%+v", c, f)
	}
}
