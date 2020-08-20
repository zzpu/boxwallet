package util

import "github.com/satori/go.uuid"

func GetUUid() string {
	return uuid.NewV4().String()
}
