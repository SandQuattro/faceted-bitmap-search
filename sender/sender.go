package sender

import (
	"errors"
	"go-bitmask-search/searcher"
)

// SendMessage sends message
func SendMessage(user, options uint32, message string) error {
	if options&searcher.SendSMS != 0 {
		return SendSMS(user, message)
	}
	if options&searcher.SendEmail != 0 {
		return SendEmail(user, message)
	}
	if options&searcher.SendPushInApp != 0 {
		return SendInApp(user, message)
	}
	if options&searcher.SendTelegram != 0 {
		return SendInTelegram(user, message)
	}
	if options&searcher.SendDiscord != 0 {
		return SendInDiscord(user, message)
	}
	if options&searcher.SendSlack != 0 {
		return SendInSlack(user, message)
	}
	if options&searcher.SendMattermost != 0 {
		return SendInMattermost(user, message)
	}
	return errors.New("unknown option")
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
	//log.Printf("Sending message %s to user: %b by %s", message, user, channel)
}
