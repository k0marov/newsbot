package service

import (
	"log"
	"sync"
)

type AuthService struct {
	correctPass     string
	authenticated   map[string]struct{} // TODO: save this persistently
	authenticatedMu sync.RWMutex
}

func NewAuthService(correctPass string) *AuthService {
	return &AuthService{
		correctPass:   correctPass,
		authenticated: make(map[string]struct{}, 100),
	}
}

func (s *AuthService) PasswordEntered(chatID, pass string) (ok bool, err error) {
	if pass == s.correctPass {
		s.authenticatedMu.Lock()
		defer s.authenticatedMu.Unlock()
		s.authenticated[chatID] = struct{}{}
		log.Printf("added chat %s to list of authenticated\n", chatID)
		return true, nil
	}
	return false, nil
}

func (s *AuthService) GetAuthenticated() (chatIDs []string, err error) {
	s.authenticatedMu.RLock()
	defer s.authenticatedMu.RUnlock()
	chatIDs = make([]string, 0, len(s.authenticated))
	for chatID := range s.authenticated {
		chatIDs = append(chatIDs, chatID)
	}
	log.Printf("%v: %v\n", len(chatIDs), chatIDs)
	return chatIDs, nil
}
