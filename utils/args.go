package utils

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
)

var RandUUIDFlag = "random UUID"
var RandStringFlag = "random string"

func ParseRandomUUID(entry string) (string, error) {
	if entry == RandUUIDFlag {
		uid, err := uuid.NewUUID()
		if err != nil {
			return "", err
		}

		return uid.String(), nil
	}

	return entry, nil
}

func ParseRandomString(entry string, length uint) string {
	if entry == RandStringFlag {
		return RandomString(NumAndLettersSet, length)
	}

	return entry
}

func ParseEmail(entry string) (string, error) {
	if entry == "" {
		return "", errors.New("email is required")
	}

	_, err := mail.ParseAddress(entry)
	if err != nil {
		return "", err
	}

	return entry, nil
}
