package models

type UserSession struct {
    User string
    SessionKey string
    LastLogin string
}

func (u *UserSession) class() string {
    return "UserSessions"
}

func (u *UserSession) getKey() string {
    return "sessionkey"
}

func (u *UserSession) getValue() interface{} {
    return u.SessionKey
}

func (u *UserSession) validate() error {
    return nil
}

