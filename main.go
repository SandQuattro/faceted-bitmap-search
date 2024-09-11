package main

import (
	"fmt"
	"go-bitmask-search/searcher"
	"go-bitmask-search/sender"
	"go-bitmask-search/util"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const totalUsers = 100_000_000

func main() {
	cores := runtime.NumCPU()

	// assume that we have %totalUsers% with notification flags, just some of them armed
	// we can have int32 = 32 flags or even more, using int64
	users := make([]uint32, totalUsers)
	for i := 0; i < totalUsers; i++ {
		users[i] = rand.Uint32()
	}

	// default little endian bit order
	bitmask := searcher.CreateBitmask(searcher.SendSlack | searcher.SendSMS | searcher.SendTelegram)
	fmt.Print("Created Bitmask: ")
	util.PrintAsBinary(bitmask)

	g := errgroup.Group{}

	// STAGE 1, BINARY SEARCH
	result := stage1(&g, users, bitmask, cores)

	// STAGE 2, SENDING MESSAGES, USING PREVIOUS STEP CHANNELS FOUND
	stage2(&g, bitmask, result)

}

func stage1(g *errgroup.Group, users []uint32, bitmask uint32, cores int) []uint32 {
	result := make([]uint32, 0)
	chunkSize := len(users) / cores

	mu := sync.Mutex{}
	once := sync.Once{}
	var start time.Time

	for i := 0; i < cores; i++ {
		g.Go(func() error {
			once.Do(func() {
				log.Print("time")
				start = time.Now()
			})
			s := 0
			e := chunkSize

			if i > 0 {
				s = chunkSize * i
				e = s + chunkSize
				if i == cores-1 {
					e = len(users)
				}
			}

			log.Printf("Start searching in users subset from %d to %d, chunk: %d using bitmask: %32b\n", s, e, i, bitmask)
			found := searcher.Search(users[s:e], bitmask)

			mu.Lock()
			result = append(result, found...)
			mu.Unlock()

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d options in %d users, time: %v", len(result), len(users), time.Since(start))

	return result
}

func stage2(g *errgroup.Group, bitmask uint32, bitmaskFound []uint32) {
	cnt := 0
	// limit our concurrent processing using errgroup semaphore
	g.SetLimit(runtime.NumCPU())

	log.Printf("sending %d notifications to users\n", len(bitmaskFound))

	start := time.Now()
	for _, user := range bitmaskFound {
		cnt++
		message := "Yo! test message"

		g.Go(func() error {
			err := sender.SendMessage(user, bitmask, message)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}

	log.Printf("notifications processing elapsed: %v\ntotal notifications: %d\n", time.Since(start), len(bitmaskFound))
}
