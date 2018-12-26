package main

import (
	"os"

	"github.com/h3poteto/yuno/listener"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	l := listener.NewListener(token)
	l.Listen()
}
