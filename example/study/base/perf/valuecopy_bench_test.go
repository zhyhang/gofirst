package perf

import "testing"

type S struct{a, b, c, d, e int64}
var sX, sY, sZ = make([]S, 1000), make([]S, 1000), make([]S, 1000)
var sumX, sumY, sumZ int64

func Benchmark_Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumX = 0
		for j := 0; j < len(sX); j++ {
			sumX += sX[j].a
		}
	}
}

func Benchmark_Range_OneIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumZ = 0
		for j := range sY {
			sumZ += sY[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumY = 0
		for _, v := range sY {
			sumY += v.a
		}
	}
}