package frontend

import (
	"github.com/k0marov/newsbot/internal/frontend/router"
	"github.com/k0marov/newsbot/internal/frontend/texts"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func StartBot(token string, svc router.Service) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, context tele.Context) {
			context.Reply(texts.UnknownError)
			log.Println("ERROR:", err)
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	router := router.NewRouter(svc)

	router.DefineRoutes(b)

	b.Start()
}
