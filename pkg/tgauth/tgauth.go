package tgauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"
)

type Auth struct {
	AuthDate  int64  `json:"auth_date"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Hash      string `json:"hash"`
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url"`
}

func (a Auth) IsValid() bool {
	key := sha256.New()
	io.WriteString(key, os.Getenv("TGBOT_TOKEN"))

	hm := hmac.New(sha256.New, key.Sum(nil))
	io.WriteString(hm, a.CheckString())

	fmt.Println(hex.EncodeToString(hm.Sum(nil)), a.Hash)
	return hex.EncodeToString(hm.Sum(nil)) == a.Hash
}

func (a Auth) IsOutdated() bool {
	return (time.Now().Unix() - a.AuthDate) > 86400
}

func (a Auth) CheckString() string {
	s := fmt.Sprintf(
		"auth_date=%v\nfirst_name=%s\nid=%v",
		a.AuthDate,
		a.FirstName,
		a.Id,
	)

	if len(a.LastName) > 0 {
		s = fmt.Sprint(s, "\nlast_name=", a.LastName)
	}

	if len(a.PhotoURL) > 0 {
		s = fmt.Sprint(s, "\nphoto_url=", a.PhotoURL)
	}

	if len(a.Username) > 0 {
		s = fmt.Sprint(s, "\nusername=", a.Username)
	}

	return s
}
