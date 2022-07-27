package p00279

import (
	"math"
	"testing"
)

func BenchmarkBruteForce(b *testing.B) {
	var ans float64
	for i := 0; i < b.N; i++ {
		ans = float64(i * i)
		_ = ans
	}
}

func BenchmarkMathPow(b *testing.B) {
	var ans float64
	for i := 0; i < b.N; i++ {
		ans = math.Pow(float64(i), 2)
		_ = ans
	}
}

func BenchmarkMathSqrt(b *testing.B) {
	var ans float64
	for i := 0; i < b.N; i++ {
		ans = math.Sqrt(float64(i))
		_ = ans
	}
}
