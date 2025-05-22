package models

import (
	"gorm.io/gorm"
	"math/rand"
)

type User struct {
	gorm.Model
	PhoneNumber string
	SessionID   string
	Orders      []Order
}

func (u *User) GenerateSessionID() {
	u.SessionID = randStringRunes(12)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
