package gosort

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var testSeed rand.Source

func TestMain(m *testing.M) {
	testSeed = rand.NewSource(time.Now().UnixNano())
	os.Exit(m.Run())
}
