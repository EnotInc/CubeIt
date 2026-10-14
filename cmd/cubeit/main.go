package main

import (
	"github.com/enotinc/cubeit/internal/cube"
	renderer "github.com/enotinc/cubeit/internal/render"
)

func main() {
	c := cube.New()
	c.Scamble()
	renderer.Render(c)
}
