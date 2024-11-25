package services

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateID() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(100000)
}
