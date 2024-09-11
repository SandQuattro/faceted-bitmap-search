package searcher

import (
	"log"
	"time"
)

const (
	SendSMS        uint32 = 1 << iota // 1 << 0 = 1
	SendEmail                         // 1 << 1 = 2
	SendTelegram                      // 1 << 2 = 4
	SendDiscord                       // 1 << 3 = 8
	SendSlack                         // 1 << 4 = 16
	SendMattermost                    // 1 << 5 = 32
	SendPushInApp                     // 1 << 6 = 64
)

func CreateBitmask(option ...uint32) uint32 {
	var result uint32
	for _, val := range option {
		result |= val
	}
	return result
}

// Search options in user data lake
func Search(users []uint32, bitmask uint32) []uint32 {
	found := make([]uint32, 0, len(users)/2)
	start := time.Now()

	for _, user := range users {
		// binary.LittleEndian.PutUint32(bs, bitmask)

		// here we are searching users with specific flags
		if user&bitmask == bitmask {
			found = append(found, user)
		}
	}

	log.Printf("time elapsed: %v\ntotal users found %d\n\n", time.Since(start), len(found))
	return found
}
