package service

type AuthService struct {
	correctPass string
}

func NewAuthService(correctPass string) *AuthService {
	return &AuthService{
		correctPass: correctPass,
	}
}

func (s *AuthService) PasswordEntered(userID, pass string) (ok bool, err error) {
	if pass == s.correctPass {
		// TODO: mark userID as authenticated
		return true, nil
	}
	return false, nil
}

func (s *AuthService) GetAuthenticated() (chatIDs []string, err error) {
	// TODO: implement me
	return nil, nil
}
