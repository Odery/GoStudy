package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

//Declare the telebot token
var token string

//get the token
func init(){
	token = os.Getenv("BOT_TOKEN")
}
func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", startEndpoint)
	b.Handle("myID", idEndpoint)
	b.Handle(tele.OnText, textEndpoint)

	b.Start()
}

//Starts the bot
func startEndpoint(c tele.Context) error {
	return c.Send("You are now connected")
}

//Prints user ID
func idEndpoint(c tele.Context) error {
	return c.Send(c.Sender().ID)
}

//Replys the same message
func textEndpoint(c tele.Context) error {
	c.Send(c.Text())
}