package main

import (
	"errors"
	"fmt"
	"go-bitmask-search/searcher"
	"go-bitmask-search/sender"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"runtime"
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

const totalUsers = 100_000_000

func main() {
	// assume that we have %totalUsers% with notification flags, just some of them armed
	// we can have int32 = 32 flags or even more, using int64
	users := make([]uint32, totalUsers)
	for i := 0; i < totalUsers; i++ {
		users[i] = rand.Uint32()
	}

	bitmask := createBitmask(SendSlack)
	found := searcher.Search(users, bitmask)

	g := errgroup.Group{}

	cnt := 0
	// limit our concurrent processing using errgroup semaphore
	g.SetLimit(runtime.NumCPU())

	fmt.Printf("sending %d notifications to users\n", len(found))

	start := time.Now()
	for _, user := range found {
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

	if err := g.Wait(); err != nil {
		panic(err)
	}

	fmt.Printf("notifications processing elapsed: %v\ntotal notifications: %d\n", time.Since(start), len(found))
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
