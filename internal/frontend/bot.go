package frontend

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func StartBot(token string, svc Service) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, context tele.Context) {
			context.Reply("Oops, some error happened!")
			log.Println("ERROR:", err)
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	router := NewRouter(svc)

	router.DefineRoutes(b)

	b.Start()
}
