package listener

import (
	"fmt"
	"github.com/k0marov/newsbot/internal/core/domain"
	"gopkg.in/telebot.v3"
	"log"
	"strconv"
)

type AuthService interface {
	GetAuthenticated() (chatIDs []string, err error)
}

type Listener struct {
	b   *telebot.Bot
	ch  <-chan domain.NewsEntry
	svc AuthService
}

func NewListener(b *telebot.Bot, ch <-chan domain.NewsEntry, svc AuthService) *Listener {
	return &Listener{b: b, ch: ch, svc: svc}
}

func (l *Listener) ListenForNews() {
	log.Println("listening for news...")
	for message := range l.ch {
		authenticatedChatIDs, err := l.svc.GetAuthenticated()
		if err != nil {
			log.Println("ERROR:", fmt.Errorf("while getting authenticated users: %w", err))
		}
		log.Println("got news message, sending it to", len(authenticatedChatIDs), "chats")
		for _, chatID := range authenticatedChatIDs {
			chatIDInt, err := strconv.ParseInt(chatID, 10, 0)
			if err != nil {
				log.Panicf("parsing chat id %q as int: %v", chatID, err)
			}
			if _, err := l.b.Send(telebot.ChatID(chatIDInt), message.URL); err != nil {
				log.Println("ERROR:", fmt.Errorf("failed sending news message to %q: %w", chatID, err))
			}
		}
	}
}
