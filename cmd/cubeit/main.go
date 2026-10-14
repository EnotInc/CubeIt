package main

import (
	"fmt"
	"os"

	"github.com/enotinc/cubeit/internal/cube"
	renderer "github.com/enotinc/cubeit/internal/render"
)

func main() {
	args := os.Args
	switch len(args) {
	case 1:
		c := cube.New()
		err := c.Scamble()
		if err != nil {
			fmt.Printf("Incorrect scramble:\n\t%s", err)
			return
		}
		renderer.Render(c)

	case 3:
		if args[1] == "-l" || args[1] == "--load" {
			c := cube.New()
			err := c.ScambleFrom(args[2])
			if err != nil {
				fmt.Printf("Incorrect scramble:\n\t%s", err)
				return
			}
			renderer.Render(c)
		}

	default:
		fmt.Printf("wrong amount of arguments: %d", len(args))
	}
}
