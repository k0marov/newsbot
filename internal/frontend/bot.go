package frontend

import (
	"github.com/k0marov/newsbot/internal/frontend/listener"
	"github.com/k0marov/newsbot/internal/frontend/router"
	"github.com/k0marov/newsbot/internal/frontend/texts"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func StartBot(token string, news <-chan string, passSVC router.Service, authSVC listener.AuthService) {
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

	r := router.NewRouter(passSVC)
	r.DefineRoutes(b)

	newsListener := listener.NewListener(b, news, authSVC)
	go newsListener.ListenForNews()
	b.Start()
}
