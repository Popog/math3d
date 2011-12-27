/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "fmt"

type Vector2 [2]float32

func MakeVector2(v ...float32) (r Vector2) {
	copy(r[:], v)
	return
}

// Entrywise addition
func (v1 Vector2) Add(v2 Vector2) Vector2 {
	v1.AddThis(v2)
	return v1
}
// In-place entrywise addition
func (v1 *Vector2) AddThis(v2 Vector2) {
	for i := range v1 {
		v1[i] += v2[i]
	}
}

// Entrywise subtraction
func (v1 Vector2) Sub(v2 Vector2) Vector2 {
	v1.SubThis(v2)
	return v1
}
// In-place entrywise subtraction
func (v1 *Vector2) SubThis(v2 Vector2) {
	for i := range v1 {
		v1[i] -= v2[i]
	}
}

// Entrywise product (Hadamard product?)
func (v1 Vector2) Mul(v2 Vector2) Vector2 {
	v1.MulThis(v2)
	return v1
}
// In-place entrywise product (Hadamard product?)
func (v1 *Vector2) MulThis(v2 Vector2) {
	for i := range v1 {
		v1[i] *= v2[i]
	}
}
// Entrywise quotient (Hadamard quotient?)
func (v1 Vector2) Div(v2 Vector2) Vector2 {
	v1.DivThis(v2)
	return v1
}
// In-place entrywise quotient (Hadamard quotient?)
func (v1 *Vector2) DivThis(v2 Vector2) {
	for i := range v1 {
		v1[i] /= v2[i]
	}
}

// Scalar multiplication
func (v Vector2) ScalarMultiply(scalar float32) Vector2 {
	v.ScalarMultiplyThis(scalar)
	return v
}
// In place scalar multiplication
func (v *Vector2) ScalarMultiplyThis(scalar float32) {
	for i := range v {
		v[i] *= scalar
	}
}

func (v1 Vector2) Dot(v2 Vector2) (r float32) {
	for i := range v1 {
		r += v1[i] * v2[i]
	}
	return
}

// The magnitude squared of a vector
func (v Vector2) LengthSq() (m float32) {
	return v.Dot(v)
}
// The magnitude of a vector
func (v Vector2) Length() float32 {
	return Sqrtf(v.LengthSq())
}

// If two vectors represents points the distance squared between them can be calculated
func (v0 Vector2) DistanceSq(v1 Vector2) float32 {
	return v0.Sub(v1).LengthSq()
}
// If two vectors represents points the distance between them can be calculated
func (v0 Vector2) Distance(v1 Vector2) float32 {
	return Sqrtf(v0.DistanceSq(v1))
}

// Normalize will modify this vector
func (v Vector2) Normalize() Vector2 {
	return v.ScalarMultiply(1.0 / v.Length())
}
// In place normalize
func (v *Vector2) NormalizeThis() {
	v.ScalarMultiplyThis(1.0 / v.Length())
}

func (v1 Vector2) Equals(v2 Vector2) bool {
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}

func (v1 Vector2) ApproxEquals(v2 Vector2, ε float32) bool {
	for i := range v1 {
		if !ApproxEquals(v1[i], v2[i], ε) {
			return false
		}
	}
	return true
}

func (v Vector2) String() string {
	return fmt.Sprintf("[%.5f,%.5f]", v[0], v[1])
}
