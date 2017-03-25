package models

import (
    "encoding/base64"
    "crypto/rand"
    "time"
)

type UserSession struct {
    User string
    SessionKey string
    LastLogin time.Time
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
    b, err := GenerateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b), err
}

func (u *UserSession) GenerateKey() error {
    key, err:= GenerateRandomString(32)
    if err != nil {
        return err
    }
    u.SessionKey = u.User + key
    return nil
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

