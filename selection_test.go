package gosort

import (
	"github.com/tomjonandy/gosort/pkg/popintslice"
	"testing"
)

func benchmarkSelection(size int, b *testing.B) {
	unsorted := popintslice.RandomSeeded(size, testSeed)
	tosort := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		Selection(tosort)
	}
}

func BenchmarkSelection100(b *testing.B)    { benchmarkSelection(100, b) }
func BenchmarkSelection1000(b *testing.B)   { benchmarkSelection(1000, b) }
func BenchmarkSelection10000(b *testing.B)  { benchmarkSelection(10000, b) }
func BenchmarkSelection100000(b *testing.B) { benchmarkSelection(100000, b) }

func BenchmarkSelectionOrdered1000(b *testing.B) {
	unsorted := popintslice.Ordered(1000)
	tosort := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		Selection(tosort)
	}
}

func benchmarkSelectionRecursive(size int, b *testing.B) {
	unsorted := popintslice.RandomSeeded(size, testSeed)
	tosort := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		SelectionRecursive(tosort)
	}
}

func BenchmarkSelectionRecursive100(b *testing.B)    { benchmarkSelectionRecursive(100, b) }
func BenchmarkSelectionRecursive1000(b *testing.B)   { benchmarkSelectionRecursive(1000, b) }
func BenchmarkSelectionRecursive10000(b *testing.B)  { benchmarkSelectionRecursive(10000, b) }
func BenchmarkSelectionRecursive100000(b *testing.B) { benchmarkSelectionRecursive(100000, b) }

func BenchmarkSelectionRecursiveOrdered1000(b *testing.B) {
	unsorted := popintslice.Ordered(1000)
	tosort := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		SelectionRecursive(tosort)
	}
}
