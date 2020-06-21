package main

import (
	"fmt"
	"math"
	"time"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (c circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getShapeArea(s shape) float64 {
	return s.getArea()
}

func main() {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("Execution started >", formatted)

	s1 := square{
		2,
	}
	// fmt.Println(s1.getArea())

	c1 := circle{
		3,
	}
	// fmt.Println(c1.getArea())

	fmt.Println("s1 area >", getShapeArea(s1))
	fmt.Println("c1 area >", getShapeArea(c1))

	t = time.Now()
	formatted = fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("Execution ended > ", formatted)
}
