package entity

import "fmt"

type Auth struct {
	AuthDate  int64  `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	Id        int    `json:"id"`
	Username  string `json:"username"`

	Token string `json:"token,omitempty"`
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
