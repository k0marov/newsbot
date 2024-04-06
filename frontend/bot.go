package frontend

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func StartBot() {
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	router := NewRouter()

	router.DefineRoutes(b)

	b.Start()
}
