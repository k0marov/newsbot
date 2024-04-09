package frontend

import (
	"github.com/k0marov/newsbot/internal/core/domain"
	"github.com/k0marov/newsbot/internal/frontend/listener"
	"github.com/k0marov/newsbot/internal/frontend/router"
	"github.com/k0marov/newsbot/internal/frontend/texts"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func StartBot(token string, news <-chan domain.NewsEntry, passSVC router.Service, authSVC listener.AuthService) (stop func()) {
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
		log.Fatalf("starting bot: %v", err)
	}

	r := router.NewRouter(passSVC)
	r.DefineRoutes(b)

	newsListener := listener.NewListener(b, news, authSVC)
	go newsListener.ListenForNews()
	go b.Start()

	return b.Stop
}
