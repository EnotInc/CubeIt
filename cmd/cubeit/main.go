package main

import (
	"github.com/enotinc/cubeit/internal/cube"
	renderer "github.com/enotinc/cubeit/internal/render"
)

func main() {
	c := cube.New()
	renderer.Render(c)
}
