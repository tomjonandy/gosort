package gosort

import (
	"github.com/tomjonandy/gosort/pkg/popintslice"
	"testing"
)

func benchmarkQuick(size int, b *testing.B) {
	unsorted := popintslice.RandomSeeded(size, testSeed)
	tosort := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		Quick(tosort)
	}
}

func BenchmarkQuick100(b *testing.B)    { benchmarkQuick(100, b) }
func BenchmarkQuick1000(b *testing.B)   { benchmarkQuick(1000, b) }
func BenchmarkQuick10000(b *testing.B)  { benchmarkQuick(10000, b) }
func BenchmarkQuick100000(b *testing.B) { benchmarkQuick(100000, b) }

func BenchmarkQuickOrdered1000(b *testing.B) {
	unsorted := popintslice.Ordered(1000)
	tosort := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		Quick(tosort)
	}
}

func benchmarkQuickConc(size int, b *testing.B) {
	unsorted := popintslice.RandomSeeded(size, testSeed)
	tosort := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		QuickConc(tosort)
	}
}

func BenchmarkQuickConc100(b *testing.B)    { benchmarkQuickConc(100, b) }
func BenchmarkQuickConc1000(b *testing.B)   { benchmarkQuickConc(1000, b) }
func BenchmarkQuickConc10000(b *testing.B)  { benchmarkQuickConc(10000, b) }
func BenchmarkQuickConc100000(b *testing.B) { benchmarkQuickConc(100000, b) }

func BenchmarkQuickConcOrdered1000(b *testing.B) {
	unsorted := popintslice.Ordered(1000)
	tosort := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		QuickConc(tosort)
	}
}
