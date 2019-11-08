package listener

import (
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

type Listener struct {
	client *slack.Client
	user   *userData
}

type userData struct {
	ID   string
	Name string
}

func NewListener(token string) *Listener {
	client := slack.New(token)
	return &Listener{
		client: client,
		user:   nil,
	}
}

func (l *Listener) Listen() {
	rtm := l.client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello
		case *slack.ConnectedEvent:
			log.Infof("Infos: %+v", *ev.Info)
			log.Info("Connection counter:", ev.ConnectionCount)

			// Update user data from connected event, to get bot user name and ID.
			user := &userData{
				ID:   ev.Info.User.ID,
				Name: ev.Info.User.Name,
			}
			l.user = user

		case *slack.MessageEvent:
			log.Debugf("event: %+v", ev)
			err := l.MessageHandler(ev, rtm)
			if err != nil {
				log.Error(err)
			}

		case *slack.PresenceChangeEvent:
			log.Infof("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			// Ignore latency
		case *slack.RTMError:
			log.Errorf("Error: %s", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Info("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
