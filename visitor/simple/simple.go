package simple

import "fmt"

const pi = 3.1415926

// Shape 元件
type Shape interface {
	getType() string
	accept(Visitor)
}

// Square 具体元件
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Circle 具体元件
type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Rectangle 具体元件
type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

// Visitor 访问者
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// AreaCalculator 具体访问者
type AreaCalculator struct {
	area any
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	a.area = s.side * s.side
	fmt.Println("square area = ", a.area)
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	a.area = pi * s.radius * s.radius
	fmt.Println("circle area = ", a.area)
}
func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	a.area = s.b * s.l
	fmt.Println("rectangle area = ", a.area)
}

// MiddleCoordinates 具体访问者
type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
