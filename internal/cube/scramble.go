package cube

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func (c *Cube) Scamble() error {
	c.Scramble = getRandomString()
	err := c.performScramble()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cube) ScambleFrom(s string) error {
	c.Scramble = s
	err := c.performScramble()
	if err != nil {
		return err
	}
	return nil
}

func getRandomString() string {
	n := 15 + rand.IntN(10)

	var scramble strings.Builder
	var prev string = "_"
	var move string = "_"

	for range n {
		for prev[0] == move[0] {
			mIndex := rand.IntN(len(moves))
			move = moves[mIndex]
		}

		scramble.WriteString(move)
		scramble.WriteString(" ")

		prev = move
	}

	return scramble.String()
}

var moves = []string{"U", "U2", "U'", "D", "D2", "D'", "F", "F2", "F'", "B", "B2", "B'", "L", "L2", "L'", "R", "R2", "R'"}

func (c *Cube) performScramble() error {
	scramble := strings.TrimSpace(c.Scramble)
	_moves := strings.SplitSeq(scramble, " ")

	for m := range _moves {
		switch m {

		case "U":
			c.rotate(U)
		case "U2":
			c.rotate(U)
			c.rotate(U)
		case "U'":
			c.rotate(Up)

		case "D":
			c.rotate(D)
		case "D2":
			c.rotate(D)
			c.rotate(D)
		case "D'":
			c.rotate(Dp)

		case "F":
			c.rotate(F)
		case "F2":
			c.rotate(F)
			c.rotate(F)
		case "F'":
			c.rotate(Fp)

		case "B":
			c.rotate(B)
		case "B2":
			c.rotate(B)
			c.rotate(B)
		case "B'":
			c.rotate(Bp)

		case "L":
			c.rotate(L)
		case "L2":
			c.rotate(L)
			c.rotate(L)
		case "L'":
			c.rotate(Lp)

		case "R":
			c.rotate(R)
		case "R2":
			c.rotate(R)
			c.rotate(R)
		case "R'":
			c.rotate(Rp)

		default:
			return fmt.Errorf("unknown move: %s", m)
		}
	}
	return nil
}
