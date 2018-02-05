package gosort

import (
	"github.com/tomjonandy/gosort/pkg/popintslice"
	"testing"
)

func benchmarkBubble(size int, b *testing.B) {
	unsorted := popintslice.RandomSeeded(size, testSeed)
	tosort := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		Bubble(tosort)
	}
}

func BenchmarkBubble100(b *testing.B)    { benchmarkBubble(100, b) }
func BenchmarkBubble1000(b *testing.B)   { benchmarkBubble(1000, b) }
func BenchmarkBubble10000(b *testing.B)  { benchmarkBubble(10000, b) }
func BenchmarkBubble100000(b *testing.B) { benchmarkBubble(100000, b) }

func BenchmarkBubbleOrdered1000(b *testing.B) {
	unsorted := popintslice.Ordered(1000)
	tosort := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		Bubble(tosort)
	}
}

func benchmarkBetterBubble(size int, b *testing.B) {
	unsorted := popintslice.RandomSeeded(size, testSeed)
	tosort := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		BetterBubble(tosort)
	}
}

func BenchmarkBetterBubble100(b *testing.B)    { benchmarkBetterBubble(100, b) }
func BenchmarkBetterBubble1000(b *testing.B)   { benchmarkBetterBubble(1000, b) }
func BenchmarkBetterBubble10000(b *testing.B)  { benchmarkBetterBubble(10000, b) }
func BenchmarkBetterBubble100000(b *testing.B) { benchmarkBetterBubble(100000, b) }

func BenchmarkBetterBubbleOrdered1000(b *testing.B) {
	unsorted := popintslice.Ordered(1000)
	tosort := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(tosort, unsorted)
		BetterBubble(tosort)
	}
}
