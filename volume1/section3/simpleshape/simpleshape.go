// Package simpleshape is a simple interface for representing geometric shapes.
package simpleshape

// The Shape type represents a geometric shape.
type Shape interface {
	Area() float64
}

// Calculates the area of a shape. The area calculation is based on the type of shape s.
func ShapeArea(s Shape) float64 {

	return s.Area()

}
