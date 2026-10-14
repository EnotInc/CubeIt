package cube

import (
	"testing"
)

var allYellow = [9]color{yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow}
var allWhite = [9]color{white, white, white, white, white, white, white, white, white}
var allGreen = [9]color{green, green, green, green, green, green, green, green, green}
var allOrange = [9]color{orange, orange, orange, orange, orange, orange, orange, orange, orange}
var allBlue = [9]color{blue, blue, blue, blue, blue, blue, blue, blue, blue}
var allRed = [9]color{red, red, red, red, red, red, red, red, red}

func testCube(t *testing.T) *Cube {
	c := New()
	testDefaultCube(t, c)
	return c
}

func testDefaultCube(t *testing.T, c *Cube) {
	testFace(t, c.faces[up], allWhite)
	testFace(t, c.faces[down], allYellow)
	testFace(t, c.faces[front], allGreen)
	testFace(t, c.faces[back], allBlue)
	testFace(t, c.faces[left], allOrange)
	testFace(t, c.faces[right], allRed)
}

func TestCubeUMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(U)
	if err != nil {
		t.Errorf("Rotation error: %s", err)
	}

	testFace(t, c.faces[up], [9]color{white, white, white, white, white, white, white, white, white})
	testFace(t, c.faces[front], [9]color{red, red, red, green, green, green, green, green, green})
	testFace(t, c.faces[right], [9]color{blue, blue, blue, red, red, red, red, red, red})
	testFace(t, c.faces[back], [9]color{orange, orange, orange, blue, blue, blue, blue, blue, blue})
	testFace(t, c.faces[left], [9]color{green, green, green, orange, orange, orange, orange, orange, orange})

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

	testFace(t, c.faces[down], [9]color{yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow})
	testFace(t, c.faces[front], [9]color{green, green, green, green, green, green, orange, orange, orange})
	testFace(t, c.faces[right], [9]color{red, red, red, red, red, red, green, green, green})
	testFace(t, c.faces[back], [9]color{blue, blue, blue, blue, blue, blue, red, red, red})
	testFace(t, c.faces[left], [9]color{orange, orange, orange, orange, orange, orange, blue, blue, blue})

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
