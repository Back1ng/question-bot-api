package tgauth

import "errors"

var (
	TokenIsOutdated = errors.New("token is outdated")
	AuthNotValid    = errors.New("auth not valid")
)
