// Package popintslice provides pre-populated int slices
package popintslice

import (
	"math/rand"
	"time"
)

// Random creates an []int populated with random ints.
// These range from 0 to 10*size.
// It returns the created slice.
func Random(size int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(seed)
	result := make([]int, size)
	maxNum := size * 10
	for i := 0; i < size; i++ {
		result[i] = myRand.Intn(maxNum)
	}
	return result
}

// RandomSeeded creates an []int populated with random ints from a specified seed.
// These range from 0 to 10*size.
// It returns the created slice.
func RandomSeeded(size int, seed rand.Source) []int {
	myRand := rand.New(seed)
	result := make([]int, size)
	maxNum := size * 10
	for i := 0; i < size; i++ {
		result[i] = myRand.Intn(maxNum)
	}
	return result
}

// Ordered creates an []int populated with consecutive ascending ints.
// These start at 1 and go up to size.
// It returns the created slice.
func Ordered(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i + 1
	}
	return result
}

// Backwards creates an []int populated with consecutive descending ints.
// These start at size and go down to 1.
// It returns the created slice.
func Backwards(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = size - i
	}
	return result
}
