package tgauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Auth struct {
	AuthDate  string `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	Id        string `json:"id"`
	Username  string `json:"username"`
}

func (a Auth) IsValid() bool {
	key := sha256.New()
	key.Write([]byte(os.Getenv("TGBOT_TOKEN")))

	hm := hmac.New(sha256.New, key.Sum(nil))
	hm.Write([]byte(a.CheckString()))

	fmt.Println(hex.EncodeToString(hm.Sum(nil)), a.Hash)
	return hex.EncodeToString(hm.Sum(nil)) == a.Hash
}

func (a Auth) IsOutdated() bool {
	authDate, _ := strconv.Atoi(a.AuthDate)

	return (time.Now().Unix() - int64(authDate)) > 86400
}

func (a Auth) CheckString() string {
	return fmt.Sprintf(
		"auth_date=%v\nfirst_name=%s\nid=%v\nusername=%s",
		a.AuthDate,
		a.FirstName,
		a.Id,
		a.Username,
	)
}
