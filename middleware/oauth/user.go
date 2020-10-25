package oauth

import "time"

type SimpleUser struct {
	ID          string `json:"id"`
	Nickname    string `json:"nickname"`
	AccessToken string
	LoginTime   time.Time
}

func (u *SimpleUser) GetAuthorization() string {
	return u.AccessToken
}

func (u *SimpleUser) GetAuthorizedAt() time.Time {
	return u.LoginTime
}

func (u *SimpleUser) GetUsername() string {
	return u.Nickname
}

func (u *SimpleUser) GetPassword() string {
	return ""
}

func (u *SimpleUser) GetEmail() string {
	return ""
}
