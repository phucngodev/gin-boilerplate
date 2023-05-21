package uuid

import (
	"strings"

	"github.com/google/uuid"
)

func GenUUID() string {
	return uuid.New().String()
}

func GenUUID16() string {
	uuidStr := uuid.New().String()

	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	return uuidStr[0:16]
}

func ParseUUIDFromStr(str string) (string, error) {
	u, err := uuid.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
