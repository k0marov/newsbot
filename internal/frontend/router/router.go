package router

import (
	"fmt"
	"github.com/k0marov/newsbot/internal/frontend/texts"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

type Service interface {
	// PasswordEntered checks password. If it is correct, it returns true and saves user as logged in.
	// otherwise, it returns false.
	PasswordEntered(userID, pass string) (ok bool, err error)
}

type Router struct {
	svc Service
}

func NewRouter(svc Service) *Router {
	return &Router{svc: svc}
}

func (r *Router) DefineRoutes(b *tele.Bot) {
	b.Handle("/start", r.Start)
	b.Handle(tele.OnText, r.HandleText)
}

func (r *Router) Start(c tele.Context) error {
	return c.Send(texts.Start)
}

func (r *Router) HandleText(c tele.Context) error {
	pass := c.Message().Text
	ok, err := r.svc.PasswordEntered(strconv.FormatInt(c.Chat().ID, 10), pass)
	if err != nil {
		return fmt.Errorf("passing password to service: %w", err)
	}
	if ok {
		return c.Reply(texts.CorrectPassEntered)
	} else {
		return c.Reply(texts.IncorrectPassEntered)
	}
}
