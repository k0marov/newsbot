package service

import (
	"fmt"
	"log"
	"slices"
)

type Repo interface {
	AddToAuthenticated(chatID string) error
	GetAllAuthenticated() ([]string, error)
}

type AuthService struct {
	correctPass string
	repo        Repo
}

func NewAuthService(correctPass string, repo Repo) *AuthService {
	return &AuthService{
		correctPass: correctPass,
		repo:        repo,
	}
}

func (s *AuthService) PasswordEntered(chatID, pass string) (ok bool, err error) {
	if pass != s.correctPass {
		return false, nil
	}
	authenticated, err := s.GetAuthenticated()
	if err != nil {
		return true, fmt.Errorf("failed getting current authenticated list: %w", err)
	}
	if slices.Contains(authenticated, chatID) {
		log.Printf("didn't add %d to list of authenticated, because it is already there", chatID)
		return true, nil
	}

	if err := s.repo.AddToAuthenticated(chatID); err != nil {
		return true, fmt.Errorf("failed adding to authenticated list: %w", err)
	}
	log.Printf("added chat %s to list of authenticated\n", chatID)
	return true, nil
}

func (s *AuthService) GetAuthenticated() (chatIDs []string, err error) {
	authenticatedChatIDs, err := s.repo.GetAllAuthenticated()
	if err != nil {
		return nil, fmt.Errorf("calling repo: %w", err)
	}
	return authenticatedChatIDs, nil
}
