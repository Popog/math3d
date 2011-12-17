/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "math"

const internalε float32 = 0.000001
const internalεε float32 = internalε * internalε

const Rad2Deg float32 = float32(180.0 / math.Pi)
const Deg2Rad float32 = float32(math.Pi / 180.0)

// some ready converted float32 values
const Pi float32 = float32(math.Pi)
const TwoPi float32 = float32(math.Pi * 2.)
const PiHalf float32 = float32(math.Pi * .5)
const Epsilon float32 = 0.000001



// these functions only exists so that we don't have to 
// use ugly float32() and float64() convertions all over the math3d32 code 
func Sinf(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func Asinf(a float32) float32 {
	return float32(math.Asin(float64(a)))
}

func Cosf(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func Acosf(a float32) float32 {
	return float32(math.Acos(float64(a)))
}

func Tanf(a float32) float32 {
	return float32(math.Tan(float64(a)))
}
func Atanf(a float32) float32 {
	return float32(math.Atan(float64(a)))
}
func Atan2f(y, x float32) float32 {
	return float32(math.Atan2(float64(x), float64(y)))
}



func Fabsf(a float32) float32 {
	return float32(math.Abs(float64(a)))
}

// Signbit returns true if x is negative or negative zero.
func Signbit(a float32) bool {
	return math.Signbit(float64(a))
}

func Sqrtf(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func AbsMin(a, b float32) float32 {
	if Fabsf(a) < Fabsf(b) {
		return a
	}
	return b
}

func AbsMin3(a, b, c float32) float32 {
	fabsa := Fabsf(a)
	fabsb := Fabsf(b)
	fabsc := Fabsf(c)

	if fabsa < fabsb && fabsa < fabsc {
		return a
	}
	if fabsb < fabsa && fabsb < fabsc {
		return b
	}
	return c
}

func AbsMax(a, b float32) float32 {
	if Fabsf(a) > Fabsf(b) {
		return a
	}
	return b
}

// return the smallest angle between two radians
// if any of the angles are larger than -+2*Pi it won't work
func MinAngleBetween(a1, a2 float32) float32 {
	diff1 := a1 - a2
	diff2 := a1 - a2 + TwoPi
	diff3 := a1 - a2 - TwoPi

	return AbsMin3(diff1, diff2, diff3)
}

/*
func MinAngleBetweenVersion2(a1,a2 float32) float32 {
	// this solution does not care about the sign  
	var crossDiff, directDiff float32
	if a1 > a2 {
		crossDiff = TwoPi - a1 + a2
		directDiff = a1 - a2
	} else {
		crossDiff = TwoPi - a2 + a1
		directDiff = a2 - a1
	}
	if crossDiff < directDiff {
		return crossDiff
	}
	return directDiff
}
*/

/*
Tests to see if the difference between two floats exceeds ε.
*/
func ApproxEquals(f1, f2, ε float32) bool {
	return Fabsf(f1-f2) < ε
}

func ApproxEquals2(f1, f2, ε float32) bool {
	return Fabsf(f1 - f2) < ε * Max(1.0, Max(Fabsf(f1), Fabsf(f2)))
}

func AlmostEqual2sComplement(A, B float32, maxUlps uint32) bool {
	// Make sure maxUlps is non-negative and small enough that the
	// default NAN won't compare as equal to anything.
	if(maxUlps >= 1 << (mantissaSize-1)) { panic("maxUlps too big") }
	
	aInt := floatBits(A)
	// Make aInt lexicographically ordered as a twos-complement int
	if aInt < 0 { aInt = minInt - aInt }
	
	// Make bInt lexicographically ordered as a twos-complement int
	bInt := floatBits(B)
	if bInt < 0 { bInt = minInt - bInt }
	
	var intDiff int32
	if bInt < aInt { intDiff = aInt - bInt
	} else         { intDiff = bInt - aInt }
	
	return intDiff <= int32(maxUlps)
}
