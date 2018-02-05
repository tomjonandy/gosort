package main

import (
	"fmt"
	"github.com/tomjonandy/gosort"
	"github.com/tomjonandy/gosort/pkg/popintslice"
)

func main() {
	unsorted := popintslice.Random(100)
	tosort := make([]int, len(unsorted))
	copy(tosort, unsorted)
	sorted := gosort.QuickConc(tosort)
	fmt.Printf("%v\n%v\n", unsorted, sorted)
}
