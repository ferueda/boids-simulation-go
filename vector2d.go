package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

// Adds x and y of two vectors
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

// Subtracts x and y of two vectors
func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

// Multiplies x and y of two vectors
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

// Adds a value v to a Vector
func (vect Vector2D) AddV(val float64) Vector2D {
	return Vector2D{x: vect.x + val, y: vect.y + val}
}

// Multiplies a value v to a Vector
func (vect Vector2D) MultiplyV(val float64) Vector2D {
	return Vector2D{x: vect.x * val, y: vect.y * val}
}

// Dicides a value v to a Vector
func (vect Vector2D) DivideV(val float64) Vector2D {
	return Vector2D{x: vect.x / val, y: vect.y / val}
}

// Limits a vector to the passed lower and upper bounds
func (vect Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{x: math.Min(math.Max(vect.x, lower), upper), y: math.Min(math.Max(vect.y, lower), upper)}
}

// Calculates the distances between two vectors (pitagoras)
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
