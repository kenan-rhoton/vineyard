package models

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
    "time"
)

type User struct {
    Login string
    Name string
    Email string
    PasswordHash string
    IsDisabled bool
}

func Login(user string, pass string) (string, error) {
    u := &User{}
    err := Grab(u,user)
    if err != nil {
        return nil, err
    }
    s := &UserSession{User: user, LastLogin: time.Now()}
    err = s.GenerateKey()
    if err != nil {
        return nil, err
    }
    err = Insert(s)
    if err != nil {
        return nil, err
    }
    return s.SessionKey, nil
}

func (u *User) HashPassword() error {
    res, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), 16)
    if err != nil {
        return err
    }
    u.PasswordHash = string(res)
    return nil
}

func (u *User) SetPassword(in string) error {
    res, err := bcrypt.GenerateFromPassword([]byte(in), 16)
    if err != nil {
        return err
    }
    u.PasswordHash = string(res)
    return nil
}

func (u *User) CheckPassword(in string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in))
}

func (u *User) class() string {
    return "Users"
}

func (u *User) getKey() string {
    return "login"
}

func (u *User) getValue() interface{} {
    return u.Login
}

func (u *User) validate() error {
    switch {
    case len(u.Login) < 5:
        return fmt.Errorf("Username too short")
    }
    return nil
}
