package cube

import (
	"testing"
)

var allYellow = [9]Color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow}
var allWhite = [9]Color{White, White, White, White, White, White, White, White, White}
var allGreen = [9]Color{Green, Green, Green, Green, Green, Green, Green, Green, Green}
var allOrange = [9]Color{Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange}
var allBlue = [9]Color{Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue, Blue}
var allRed = [9]Color{Red, Red, Red, Red, Red, Red, Red, Red, Red}

func testCube(t *testing.T) *Cube {
	c := New()
	testDefaultCube(t, c)
	return c
}

func testFace(t *testing.T, f face, c [9]Color) {
	if f != c {
		t.Errorf("wrong face rotation.\nWant\t%+v\nGot\t%+v", c, f)
	}
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
		t.Fatalf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[UP], [9]Color{White, White, White, White, White, White, White, White, White})
	testFace(t, c.Faces[FRONT], [9]Color{Red, Red, Red, Green, Green, Green, Green, Green, Green})
	testFace(t, c.Faces[RIGHT], [9]Color{Blue, Blue, Blue, Red, Red, Red, Red, Red, Red})
	testFace(t, c.Faces[BACK], [9]Color{Orange, Orange, Orange, Blue, Blue, Blue, Blue, Blue, Blue})
	testFace(t, c.Faces[LEFT], [9]Color{Green, Green, Green, Orange, Orange, Orange, Orange, Orange, Orange})

	err = c.rotate(Up)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestCubeDMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(D)
	if err != nil {
		t.Errorf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[DOWN], [9]Color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow})
	testFace(t, c.Faces[FRONT], [9]Color{Green, Green, Green, Green, Green, Green, Orange, Orange, Orange})
	testFace(t, c.Faces[RIGHT], [9]Color{Red, Red, Red, Red, Red, Red, Green, Green, Green})
	testFace(t, c.Faces[BACK], [9]Color{Blue, Blue, Blue, Blue, Blue, Blue, Red, Red, Red})
	testFace(t, c.Faces[LEFT], [9]Color{Orange, Orange, Orange, Orange, Orange, Orange, Blue, Blue, Blue})

	err = c.rotate(Dp)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestCubeFMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(F)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[FRONT], allGreen)
	testFace(t, c.Faces[UP], [9]Color{White, White, White, White, White, White, Orange, Orange, Orange})
	testFace(t, c.Faces[LEFT], [9]Color{Orange, Orange, Yellow, Orange, Orange, Yellow, Orange, Orange, Yellow})
	testFace(t, c.Faces[DOWN], [9]Color{Red, Red, Red, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow})
	testFace(t, c.Faces[RIGHT], [9]Color{White, Red, Red, White, Red, Red, White, Red, Red})

	err = c.rotate(Fp)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestCubeBMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(B)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[BACK], allBlue)
	testFace(t, c.Faces[UP], [9]Color{Red, Red, Red, White, White, White, White, White, White})
	testFace(t, c.Faces[LEFT], [9]Color{White, Orange, Orange, White, Orange, Orange, White, Orange, Orange})
	testFace(t, c.Faces[DOWN], [9]Color{Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Orange, Orange, Orange})
	testFace(t, c.Faces[RIGHT], [9]Color{Red, Red, Yellow, Red, Red, Yellow, Red, Red, Yellow})

	err = c.rotate(Bp)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestCubeRMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(R)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[RIGHT], allRed)
	testFace(t, c.Faces[UP], [9]Color{White, White, Green, White, White, Green, White, White, Green})
	testFace(t, c.Faces[FRONT], [9]Color{Green, Green, Yellow, Green, Green, Yellow, Green, Green, Yellow})
	testFace(t, c.Faces[DOWN], [9]Color{Yellow, Yellow, Blue, Yellow, Yellow, Blue, Yellow, Yellow, Blue})
	testFace(t, c.Faces[BACK], [9]Color{White, Blue, Blue, White, Blue, Blue, White, Blue, Blue})

	err = c.rotate(Rp)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestCubeLMoves(t *testing.T) {
	c := testCube(t)
	err := c.rotate(L)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testFace(t, c.Faces[LEFT], allOrange)
	testFace(t, c.Faces[UP], [9]Color{Blue, White, White, Blue, White, White, Blue, White, White})
	testFace(t, c.Faces[FRONT], [9]Color{White, Green, Green, White, Green, Green, White, Green, Green})
	testFace(t, c.Faces[DOWN], [9]Color{Green, Yellow, Yellow, Green, Yellow, Yellow, Green, Yellow, Yellow})
	testFace(t, c.Faces[BACK], [9]Color{Blue, Blue, Yellow, Blue, Blue, Yellow, Blue, Blue, Yellow})

	err = c.rotate(Lp)
	if err != nil {
		t.Fatalf("Rotation error: %s", err)
	}

	testDefaultCube(t, c)
}

func TestRightPifPaf(t *testing.T) {
	c := testCube(t)
	c.rotate(R)
	c.rotate(U)
	c.rotate(Rp)
	c.rotate(Up)

	testFace(t, c.Faces[UP], [9]Color{White, White, Orange, White, White, Green, White, White, Green})
	testFace(t, c.Faces[FRONT], [9]Color{Green, Green, Yellow, Green, Green, White, Green, Green, Green})
	testFace(t, c.Faces[RIGHT], [9]Color{Red, Red, White, Blue, Red, Red, White, Red, Red})
	testFace(t, c.Faces[DOWN], [9]Color{Yellow, Yellow, Red, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow})
	testFace(t, c.Faces[BACK], [9]Color{Blue, Red, Red, Blue, Blue, Blue, Blue, Blue, Blue})
	testFace(t, c.Faces[LEFT], [9]Color{Blue, Orange, Orange, Orange, Orange, Orange, Orange, Orange, Orange})
}

func TestLeftPifPaf(t *testing.T) {
	c := testCube(t)
	c.rotate(Lp)
	c.rotate(Up)
	c.rotate(L)
	c.rotate(U)

	testFace(t, c.Faces[UP], [9]Color{Red, White, White, Green, White, White, Green, White, White})
	testFace(t, c.Faces[FRONT], [9]Color{Yellow, Green, Green, White, Green, Green, Green, Green, Green})
	testFace(t, c.Faces[RIGHT], [9]Color{Red, Red, Blue, Red, Red, Red, Red, Red, Red})
	testFace(t, c.Faces[DOWN], [9]Color{Orange, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow, Yellow})
	testFace(t, c.Faces[BACK], [9]Color{Orange, Orange, Blue, Blue, Blue, Blue, Blue, Blue, Blue})
	testFace(t, c.Faces[LEFT], [9]Color{White, Orange, Orange, Orange, Orange, Blue, Orange, Orange, White})
}
