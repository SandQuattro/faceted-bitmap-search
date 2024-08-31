package sender

import "log"

func SendSMS(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b by SMS", message, user)
	return nil
}

func SendEmail(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b by Email", message, user)
	return nil
}

func SendInApp(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b in App", message, user)
	return nil
}

func SendInTelegram(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b by Telegram", message, user)
	return nil
}

func SendInDiscord(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b by Discord", message, user)
	return nil
}

func SendInSlack(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b by Slack", message, user)
	return nil
}

func SendInMattermost(user uint32, message string) error {
	log.Printf("Sending message %s to user: %b by Mattermost", message, user)
	return nil
}
