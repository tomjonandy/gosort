package gosort

import (
	"math/rand"
	"sync"
	"time"
)

func Quick(tosort []int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(seed)
	quick(tosort, myRand) // call to recursive section
	return tosort
}

// Recursive section
func quick(tosort []int, myRand *rand.Rand) []int {
	if len(tosort) > 1 {
		pivot := tosort[myRand.Intn(len(tosort))]
		tosort, leftEnd, rightBeg := dnf(tosort, pivot)
		quick(tosort[:leftEnd], myRand)
		quick(tosort[rightBeg:], myRand)
	}
	return tosort
}

func QuickConc(tosort []int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	myRand := &MutexRand{rand: rand.New(seed)}
	var wg sync.WaitGroup
	wg.Add(1)
	quickConc(tosort, myRand, &wg) // call to recursive section
	wg.Wait()
	return tosort
}

func quickConc(tosort []int, myRand *MutexRand, wg *sync.WaitGroup) []int {
	defer wg.Done()
	if len(tosort) > 1 {
		pivot := tosort[myRand.Intn(len(tosort))]
		tosort, leftEnd, rightBeg := dnf(tosort, pivot)
		wg.Add(2)
		go quickConc(tosort[:leftEnd], myRand, wg)
		go quickConc(tosort[rightBeg:], myRand, wg)
	}
	return tosort
}

type MutexRand struct {
	rand *rand.Rand
	mux  sync.Mutex
}

func (r *MutexRand) Intn(size int) int {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.rand.Intn(size)
}

// Dutch National Flag
func dnf(tosort []int, pivot int) ([]int, int, int) {
	next, leftEnd := 0, 0
	rightBeg := len(tosort)
	for next < rightBeg {
		if tosort[next] < pivot {
			tosort[leftEnd], tosort[next] = tosort[next], tosort[leftEnd]
			leftEnd++
			next++
		} else if tosort[next] > pivot {
			rightBeg--
			tosort[rightBeg], tosort[next] = tosort[next], tosort[rightBeg]
		} else {
			next++
		}
	}
	return tosort, leftEnd, rightBeg
}
