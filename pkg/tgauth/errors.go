package tgauth

import "errors"

var TokenIsOutdated = errors.New("token is outdated")
var AuthNotValid = errors.New("auth not valid")
