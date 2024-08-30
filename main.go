package main

import (
	"fmt"
	"log"
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

const totalUsers = 100_000_000

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
	for _, user := range users {
		bitmask := SendEmail | SendSlack
		//binary.LittleEndian.PutUint32(bs, bitmask)
		if user&bitmask == bitmask {
			// here we are searching users with specific flags
			cnt++
			// message := "Yo! test message"
			// go SendMessage(user, options, message)
		}
	}

	fmt.Printf("time elapsed(sec): %f\ntotal users found %d\n", time.Since(start).Seconds(), cnt)
}

func SendMessage(user, options uint32, message string) {
	if options&SendSMS != 0 {
		sendSMS(user, message)
	}
	if options&SendEmail != 0 {
		sendEmail(user, message)
	}
	if options&SendPushInApp != 0 {
		sendInApp(user, message)
	}
	if options&SendTelegram != 0 {
		sendInTelegram(user, message)
	}
	if options&SendDiscord != 0 {
		sendInDiscord(user, message)
	}
	if options&SendSlack != 0 {
		//sendInSlack(user, message)
	}
	if options&SendMattermost != 0 {
		//sendInMattermost(user, message)
	}
}

func sendSMS(user uint32, message string) {
	log.Printf("Sending message %s to user: %b by SMS", message, user)
}

func sendEmail(user uint32, message string) {
	log.Printf("Sending message %s to user: %b by Email", message, user)
}

func sendInApp(user uint32, message string) {
	log.Printf("Sending message %s to user: %b in App", message, user)
}

func sendInTelegram(user uint32, message string) {
	log.Printf("Sending message %s to user: %b by Telegram", message, user)
}

func sendInDiscord(user uint32, message string) {
	log.Printf("Sending message %s to user: %b by Discord", message, user)
}

func sendInSlack(user uint32, message string) {
	log.Printf("Sending message %s to user: %b by Slack", message, user)
}

func sendInMattermost(user uint32, message string) {
	log.Printf("Sending message %s to user: %b by Mattermost", message, user)
}
