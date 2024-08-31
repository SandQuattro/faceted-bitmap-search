package searcher

import (
	"math/rand/v2"
	"testing"
)

const TOTAL = 100_000_000

func Benchmark_Search(b *testing.B) {
	b.StopTimer()

	users := make([]uint32, TOTAL)

	for i := 0; i < TOTAL; i++ {
		users[i] = rand.Uint32()
	}

	bitmask := rand.Uint32N(6)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Search(users, bitmask)
	}
}
