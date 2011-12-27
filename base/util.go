/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "math"

const internalε floatType = 0.000001
const internalεε floatType = internalε * internalε

const Rad2Deg floatType = floatType(180.0 / math.Pi)
const Deg2Rad floatType = floatType(math.Pi / 180.0)

// some ready converted floatType values
const Pi floatType = floatType(math.Pi)
const TwoPi floatType = floatType(math.Pi * 2.)
const PiHalf floatType = floatType(math.Pi * .5)
const Epsilon floatType = 0.000001

// these functions only exists so that we don't have to 
// use ugly float32() and float64() convertions all over the math3d32 code 
func Sinf(a floatType) floatType {
	return floatType(math.Sin(float64(a)))
}

func Asinf(a floatType) floatType {
	return floatType(math.Asin(float64(a)))
}

func Cosf(a floatType) floatType {
	return floatType(math.Cos(float64(a)))
}

func Acosf(a floatType) floatType {
	return floatType(math.Acos(float64(a)))
}

func Tanf(a floatType) floatType {
	return floatType(math.Tan(float64(a)))
}
func Atanf(a floatType) floatType {
	return floatType(math.Atan(float64(a)))
}
func Atan2f(y, x floatType) floatType {
	return floatType(math.Atan2(float64(x), float64(y)))
}

func Fabsf(a floatType) floatType {
	return floatType(math.Abs(float64(a)))
}

// Signbit returns true if x is negative or negative zero.
func Signbit(a floatType) bool {
	return math.Signbit(float64(a))
}

func Sqrtf(a floatType) floatType {
	return floatType(math.Sqrt(float64(a)))
}

func Min(a, b floatType) floatType {
	if a < b {
		return a
	}
	return b
}

func Max(a, b floatType) floatType {
	if a > b {
		return a
	}
	return b
}

func AbsMin(a, b floatType) floatType {
	if Fabsf(a) < Fabsf(b) {
		return a
	}
	return b
}

func AbsMin3(a, b, c floatType) floatType {
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

func AbsMax(a, b floatType) floatType {
	if Fabsf(a) > Fabsf(b) {
		return a
	}
	return b
}

// return the smallest angle between two radians
// if any of the angles are larger than -+2*Pi it won't work
func MinAngleBetween(a1, a2 floatType) floatType {
	diff1 := a1 - a2
	diff2 := a1 - a2 + TwoPi
	diff3 := a1 - a2 - TwoPi

	return AbsMin3(diff1, diff2, diff3)
}

/*
func MinAngleBetweenVersion2(a1,a2 floatType) floatType {
	// this solution does not care about the sign  
	var crossDiff, directDiff floatType
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
func ApproxEquals(f1, f2, ε floatType) bool {
	return Fabsf(f1-f2) < ε
}

func ApproxEquals2(f1, f2, ε floatType) bool {
	return Fabsf(f1-f2) < ε*Max(1.0, Max(Fabsf(f1), Fabsf(f2)))
}

func AlmostEqual2sComplement(A, B floatType, maxUlps uintType) bool {
	// Make sure maxUlps is non-negative and small enough that the
	// default NAN won't compare as equal to anything.
	if maxUlps >= 1<<(mantissaSize-1) {
		panic("maxUlps too big")
	}

	aInt := floatBits(A)
	// Make aInt lexicographically ordered as a twos-complement int
	if aInt < 0 {
		aInt = minInt - aInt
	}

	// Make bInt lexicographically ordered as a twos-complement int
	bInt := floatBits(B)
	if bInt < 0 {
		bInt = minInt - bInt
	}

	var intDiff intType
	if bInt < aInt {
		intDiff = aInt - bInt
	} else {
		intDiff = bInt - aInt
	}

	return intDiff <= intType(maxUlps)
}
