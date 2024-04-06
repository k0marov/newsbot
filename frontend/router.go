package frontend

import tele "gopkg.in/telebot.v3"

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) DefineRoutes(b *tele.Bot) {
	b.Handle("/hello", r.Hello)
}

func (r *Router) Hello(c tele.Context) error {
	return c.Send("Hello!")
}
