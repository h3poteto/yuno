package listener

import (
	"os"
	"regexp"

	"github.com/h3poteto/yuno/kingtime"
	"github.com/nlopes/slack"
)

func (l *Listener) MessageHandler(message *slack.MessageEvent, rtm *slack.RTM) error {
	start := regexp.MustCompile(`おはー`)
	if start.MatchString(message.Text) {
		client := kingtime.New("https://api.kingtime.jp/v1.0", os.Getenv("KING_OF_TIME_TOKEN"))
		_, err := client.Attendance("33001")
		if err != nil {
			return err
		}
		rtm.SendMessage(rtm.NewOutgoingMessage("おはー 打刻したよー", message.Channel))
		return nil
	}
	end := regexp.MustCompile(`店じまい`)
	if end.MatchString(message.Text) {
		// TODO: だこく
		rtm.SendMessage(rtm.NewOutgoingMessage("おつー 打刻したよー", message.Channel))
		return nil
	}
	return nil
}
