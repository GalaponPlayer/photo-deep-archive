package lib

import (
	"app/src/pkg/errorhandle"
	"time"

	"github.com/google/uuid"
)

func GenerateUUIDv4() (UUIDv4, error) {
	uuid, err := uuid.NewRandom()
	//todo: error handling
	if err != nil {
		return UUIDv4(""), errorhandle.Wrap("lib.GenerateUUIDv4()", err)
	}
	return UUIDv4(uuid.String()), nil
}

func GetNowUnixTimeSeconds() int64 {
	return time.Now().Unix()
}
