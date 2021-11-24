// Copyright 2021 The Bench Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const (
	// Size is the number bench loops
	Size = 16
)

func main() {
	bench := func() {
		a, i := 1.0, 1
		for {
			if i&1 == 1 {
				a *= float64(i)
			} else {
				a /= float64(i)
			}
			i++
		}
	}
	for i := 0; i < Size-1; i++ {
		go bench()
	}
	bench()
}
