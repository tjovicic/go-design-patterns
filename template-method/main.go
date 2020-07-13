package main

import "fmt"

type iShape interface {
	doDrawing()
}

type shape struct {
	s iShape
}

func (s *shape) draw() {
	fmt.Println("prepare graphics")
	s.s.doDrawing()
}

type circle struct{}

func (c *circle) doDrawing() {
	fmt.Println("drawing circle")
}

type line struct{}

func (l *line) doDrawing() {
	fmt.Println("drawing line")
}

func main() {
	s := shape{
		s: &circle{},
	}

	s.draw()

	s = shape{
		s: &line{},
	}

	s.draw()
}
