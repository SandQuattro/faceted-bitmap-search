package main

import (
	"errors"
	"fmt"
	"go-bitmask-search/sender"
	"golang.org/x/sync/errgroup"
	"math/rand"
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

const totalUsers = 1_000

func main() {
	// assume that we have %totalUsers% with notification flags, just some of them armed
	// we can have int32 = 32 flags or even more, using int64
	users := make([]uint32, totalUsers)
	for i := 0; i < totalUsers; i++ {
		users[i] = rand.Uint32()
	}

	cnt := 0
	//bs := make([]byte, 0)

	fmt.Printf("Start searching in users total set: %d\n", totalUsers)
	start := time.Now()

	g := errgroup.Group{}
	for _, user := range users {
		bitmask := createBitmask(SendSlack)
		//binary.LittleEndian.PutUint32(bs, bitmask)
		if user&bitmask == bitmask {
			// here we are searching users with specific flags
			cnt++
			message := "Yo! test message"

			g.Go(func() error {
				err := SendMessage(user, bitmask, message)
				if err != nil {
					return err
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}

	fmt.Printf("time elapsed: %v\ntotal users found %d\n", time.Since(start), cnt)
}

func createBitmask(option ...uint32) uint32 {
	var result uint32
	for _, val := range option {
		result |= val
	}
	return result
}

func SendMessage(user, options uint32, message string) error {
	if options&SendSMS != 0 {
		return sender.SendSMS(user, message)
	}
	if options&SendEmail != 0 {
		return sender.SendEmail(user, message)
	}
	if options&SendPushInApp != 0 {
		return sender.SendInApp(user, message)
	}
	if options&SendTelegram != 0 {
		return sender.SendInTelegram(user, message)
	}
	if options&SendDiscord != 0 {
		return sender.SendInDiscord(user, message)
	}
	if options&SendSlack != 0 {
		return sender.SendInSlack(user, message)
	}
	if options&SendMattermost != 0 {
		return sender.SendInMattermost(user, message)
	}
	return errors.New("unknown option")
}
