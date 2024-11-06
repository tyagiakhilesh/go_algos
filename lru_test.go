package algos_test

import (
	algos "algorithms"
	"math/rand"
	"testing"
)

func TestLruEviction(t *testing.T) {
	cache := algos.InitCache(10)
	count := 0
	for count < 20 {
		cache.Put(rand.Int(), rand.Int())
		count++
	}
	if cache.Size() != 10 {
		t.Errorf("Size should have been %d but found %d, "+
			"as others should have been evicted", 10, cache.Size())
	}
}
