package tgauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

type Auth struct {
	AuthDate  int64  `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	Id        int    `json:"id"`
	Username  string `json:"username"`
}

func (a Auth) IsValid() bool {
	key := sha256.New()
	key.Write([]byte(os.Getenv("TGBOT_TOKEN")))

	hm := hmac.New(sha256.New, key.Sum(nil))
	hm.Write([]byte(a.CheckString()))

	return hex.EncodeToString(hm.Sum(nil)) == a.Hash
}

func (a Auth) CheckString() string {
	return fmt.Sprintf(
		"auth_date=%d\nfirst_name=%s\nid=%d\nusername=%s",
		a.AuthDate,
		a.FirstName,
		a.Id,
		a.Username,
	)
}
