package vmath

import "math"

type Vec []float64

/* == Addition ================== */
// Add a vector to another vector
func (a Vec) Add(b Vec) Vec {
	vec := Vec{}
	for i := range a {
		vec = append(vec, a[i]+b[i])
	}

	return vec
}

// Add a number to a vector
func (a Vec) AddN(b float64) Vec {
	vec := Vec{}
	for i := range a {
		vec = append(vec, a[i]+b)
	}

	return vec
}

/* ============================== */

/* == Subtraction =============== */
// Subtract a vector from another vector
func (a Vec) Sub(b Vec) Vec {
	vec := Vec{}
	for i := range a {
		vec = append(vec, a[i]-b[i])
	}

	return vec
}

// Subtract a number from a vector
func (a Vec) SubN(b float64) Vec {
	vec := Vec{}
	for i := range a {
		vec = append(vec, a[i]-b)
	}

	return vec
}

/* ============================== */

/* == Multiplication ============ */
// Scalar multiplication for a vector
func (a Vec) ScalarMul(b float64) Vec {
	vec := Vec{}
	for i := range a {
		vec = append(vec, a[i]*b)
	}

	return vec
}

// Dot product of two vectors
func (a Vec) Dot(b Vec) float64 {
	//return (a[0] * b[0]) + (a[1] * b[1])
	product := 0.0
	for i := range a {
		product += a[i] * b[i]
	}

	return product
}

// 2D Cross product of two vectors
func (a Vec) Cross2D(b Vec) float64 {
	return (a[0] * b[1]) - (a[1] * b[0])
}

// 3D Cross product of two vectors
func (a Vec) Cross3D(b Vec) Vec {
	return Vec{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

/* ============================== */

/* == Util ====================== */
// Total dimensions of a vector
func (a Vec) D() int {
	return len(a)
}

// Length of a vector
func (a Vec) Len() float64 {
	return math.Sqrt(a.Dot(a))
}

// Normalized Vector
func (a Vec) Norm() Vec {
	return a.ScalarMul(1.0 / a.Len())
}

// ABS????
func (a Vec) Abs() Vec {
	vec := Vec{}
	for i := range a {
		vec = append(vec, math.Abs(a[i]))
	}

	return vec
}

/* ============================== */
