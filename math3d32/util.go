/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "math"
import "rand"

type Float float32

const internalε Float = 0.000001
const internalεε Float = internalε * internalε

const Rad2Deg Float = Float(180.0 / math.Pi)
const Deg2Rad Float = Float(math.Pi / 180.0)

// some ready converted Float values
const Pi Float = Float(math.Pi)
const TwoPi Float = Float(math.Pi * 2.)
const PiHalf Float = Float(math.Pi * .5)
const Epsilon Float = 0.000001


func Randf(r *rand.Rand) Float {
	return Float(rand.Float64())
}

// these functions only exists so that we don't have to 
// use ugly float32() and float64() convertions all over the math3d32 code 
func Sinf(a Float) Float {
	return Float(math.Sin(float64(a)))
}

func Asinf(a Float) Float {
	return Float(math.Asin(float64(a)))
}

func Cosf(a Float) Float {
	return Float(math.Cos(float64(a)))
}

func Acosf(a Float) Float {
	return Float(math.Acos(float64(a)))
}

func Tanf(a Float) Float {
	return Float(math.Tan(float64(a)))
}
func Atanf(a Float) Float {
	return Float(math.Atan(float64(a)))
}
func Atan2f(y, x Float) Float {
	return Float(math.Atan2(float64(x), float64(y)))
}



func Fabsf(a Float) Float {
	return Float(math.Abs(float64(a)))
}

// Signbit returns true if x is negative or negative zero.
func Signbit(a Float) bool {
	return math.Signbit(float64(a))
}

func Sqrtf(a Float) Float {
	return Float(math.Sqrt(float64(a)))
}

func Min(a, b Float) Float {
	if a < b {
		return a
	}
	return b
}

func Max(a, b Float) Float {
	if a > b {
		return a
	}
	return b
}

func AbsMin(a, b Float) Float {
	if Fabsf(a) < Fabsf(b) {
		return a
	}
	return b
}

func AbsMin3(a, b, c Float) Float {
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

func AbsMax(a, b Float) Float {
	if Fabsf(a) > Fabsf(b) {
		return a
	}
	return b
}

// return the smallest angle between two radians
// if any of the angles are larger than -+2*Pi it won't work
func MinAngleBetween(a1, a2 Float) Float {
	diff1 := a1 - a2
	diff2 := a1 - a2 + TwoPi
	diff3 := a1 - a2 - TwoPi

	return AbsMin3(diff1, diff2, diff3)
}

/*
func MinAngleBetweenVersion2(a1,a2 Float) Float {
	// this solution does not care about the sign  
	var crossDiff, directDiff Float
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
func ApproxEquals(f1, f2, ε Float) bool {
	return Fabsf(f1-f2) < ε
}

func ApproxEquals2(f1, f2, ε Float) bool {
	return Fabsf(f1 - f2) < ε * Max(1.0, Max(Fabsf(f1), Fabsf(f2)))
}

func AlmostEqual2sComplement(A, B Float, maxUlps uint) bool {
	// Make sure maxUlps is non-negative and small enough that the
	// default NAN won't compare as equal to anything.
	if(maxUlps >= 1 << (23-1)) { panic("maxUlps too big") }
	
	aInt := int(math.Float32bits(float32(A)))
	// Make aInt lexicographically ordered as a twos-complement int
	if aInt < 0 { aInt = math.MinInt32 - aInt }
	
	// Make bInt lexicographically ordered as a twos-complement int
	bInt := int(math.Float32bits(float32(B)))
	if bInt < 0 { bInt = math.MinInt32 - bInt }
	
	var intDiff int
	if bInt < aInt { intDiff = aInt - bInt
	} else         { intDiff = bInt - aInt }
	
	return intDiff <= int(maxUlps)
}
