package sender

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
