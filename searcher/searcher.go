package searcher

import (
	"log"
	"time"
)

// Search options in user data lake
func Search(users []uint32, bitmask uint32) []uint32 {
	found := make([]uint32, len(users)/2)

	log.Printf("Start searching in users total set: %d using bitmask: %32b\n", len(users), bitmask)
	start := time.Now()

	for _, user := range users {
		// binary.LittleEndian.PutUint32(bs, bitmask)
		if user&bitmask == bitmask {
			// here we are searching users with specific flags
			found = append(found, user)
		}
	}

	log.Printf("time elapsed: %v\ntotal users found %d\n\n", time.Since(start), len(found))
	return found
}
