package backend

import "errors"

type FakeService struct {
}

func (s *FakeService) PasswordEntered(userID, pass string) (ok bool, err error) {
	if pass == "hello world" {
		return true, nil
	} else if pass == "error" {
		return false, errors.New("fake error")
	}
	return false, nil
}
