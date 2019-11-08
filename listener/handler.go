package listener

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/h3poteto/yuno/kingtime"
	"github.com/h3poteto/yuno/models"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

func (l *Listener) MessageHandler(message *slack.MessageEvent, rtm *slack.RTM) error {
	if strings.HasPrefix(message.Msg.Text, fmt.Sprintf("<@%s> ", l.user.ID)) {
		return l.HandleReply(message, rtm)
	}
	return l.HandleNoReply(message, rtm)
}

func (l *Listener) HandleReply(message *slack.MessageEvent, rtm *slack.RTM) error {
	log.Infof("Message Received: %s", message.Text)
	m := strings.Split(strings.TrimSpace(message.Text), " ")[1:]
	if m[0] == "help" {
		rtm.SendMessage(rtm.NewOutgoingMessage("help - このメニューを表示\nregistration - 従業員コードを登録", message.Channel))
		return nil
	}
	if m[0] == "registration" {
		if len(m) < 2 || len(m[1]) == 0 {
			rtm.SendMessage(rtm.NewOutgoingMessage("registration 1234 のように入力してくださいね", message.Channel))
			return errors.New("Invalid code")
		}
		client := kingtime.New(os.Getenv("KING_OF_TIME_TOKEN"))
		response, err := client.GetEmployee(m[1])
		if err != nil {
			return err
		}
		log.Infof("employee: %+v", *response)
		emp, err := models.NewEmployee(response.Code, response.LastName, response.FirstName, response.Key, response.TypeName, message.Msg.User)
		if err != nil {
			return err
		}
		return emp.Save()
	}
	return nil
}

func (l *Listener) HandleNoReply(message *slack.MessageEvent, rtm *slack.RTM) error {
	start := regexp.MustCompile(`おはー`)
	if start.MatchString(message.Text) {
		client := kingtime.New(os.Getenv("KING_OF_TIME_TOKEN"))
		_, err := client.Attendance("3001")
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
