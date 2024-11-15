package sender

import (
	"errors"
	"go-bitmask-search/searcher"
	"log"
)

// SendMessage sends message
func SendMessage(user, options uint32, message string) error {
	var (
		err   error
		found bool
	)

	if options&searcher.SendSMS != 0 {
		err = SendSMS(user, message)
		found = true
	}
	if options&searcher.SendEmail != 0 {
		err = SendEmail(user, message)
		found = true
	}
	if options&searcher.SendPushInApp != 0 {
		err = SendInApp(user, message)
		found = true
	}
	if options&searcher.SendTelegram != 0 {
		err = SendInTelegram(user, message)
		found = true
	}
	if options&searcher.SendDiscord != 0 {
		err = SendInDiscord(user, message)
		found = true
	}
	if options&searcher.SendSlack != 0 {
		err = SendInSlack(user, message)
		found = true
	}
	if options&searcher.SendMattermost != 0 {
		err = SendInMattermost(user, message)
		found = true
	}
	if !found {
		return errors.New("options not fond")
	}

	return err
}

func SendSMS(user uint32, message string) error {
	helper(user, message, "SMS")
	return nil
}

func SendEmail(user uint32, message string) error {
	helper(user, message, "Email")
	return nil
}

func SendInApp(user uint32, message string) error {
	helper(user, message, "InApp")
	return nil
}

func SendInTelegram(user uint32, message string) error {
	helper(user, message, "Telegram")
	return nil
}

func SendInDiscord(user uint32, message string) error {
	helper(user, message, "Discord")
	return nil
}

func SendInSlack(user uint32, message string) error {
	helper(user, message, "Slack")
	return nil
}

func SendInMattermost(user uint32, message string) error {
	helper(user, message, "Mattermost")
	return nil
}

func helper(user uint32, message, channel string) {
	log.Printf("Sending message %s to user: %b by %s", message, user, channel)
}
