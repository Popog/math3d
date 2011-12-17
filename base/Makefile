# Copyright 2010 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include base_Makefile

generate: clean
	cp *.go ../math3d32
	cp helpers.go.32 ../math3d32/helpers.go
	cp base_Makefile ../math3d32/Makefile
	sed -i 's/floatType/float32/g;s/uintType/uint32/g;s/intType/int32/g' ../math3d32/*.go
	cp *.go ../math3d64
	cp helpers.go.64 ../math3d64/helpers.go
	cp base_Makefile ../math3d64/Makefile
	sed -i 's/floatType/float64/g;s/uintType/uint64/g;s/intType/int64/g;s/math3d32/math3d64/g' ../math3d64/*.go
	sed -i 's/math3d32/math3d64/g' ../math3d64/Makefile
