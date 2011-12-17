/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d64

import "testing"

func BenchmarkDummy(b *testing.B) {
	// b.StopTimer()
	// const size = 3
	// r := rand.New(rand.NewSource(time.Nanoseconds()))
	// data := make([]float32, size*size)
	
	// b.StartTimer()
	var count int
	for i := 0; i < b.N; i++ {
		count++
	}
}

func BenchmarkApproxEquals(b *testing.B) {
	// b.StopTimer()
	// const size = 3
	// r := rand.New(rand.NewSource(time.Nanoseconds()))
	// data := make([]float32, size*size)
	
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		ApproxEquals(1, 1.5, 0.5)
	}
}

func BenchmarkApproxEquals2(b *testing.B) {
	// b.StopTimer()
	// const size = 3
	// r := rand.New(rand.NewSource(time.Nanoseconds()))
	// data := make([]float32, size*size)
	
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		ApproxEquals2(1, 1.5, 0.5)
	}
}

func BenchmarkAlmostEqual2sComplement(b *testing.B) {
	// b.StopTimer()
	// const size = 3
	// r := rand.New(rand.NewSource(time.Nanoseconds()))
	// data := make([]float32, size*size)
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		AlmostEqual2sComplement(1, 1.5, 500)
	}
}