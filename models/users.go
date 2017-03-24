package models

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

type User struct {
    Login string
    Name string
    PasswordHash string
    IsDisabled bool
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
