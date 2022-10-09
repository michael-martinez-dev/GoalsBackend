package util

import (
	"math/rand"
	"strings"
	"time"
)

type JError struct {
	Error string `json:"error"`
}

func NewJError(err error) JError {
	jerr := JError{"generic error"}
	if err != nil {
		jerr.Error = err.Error()
	}
	return jerr
}

func GenerateId() string {
	return strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(GenerateRandomString(20))), " ", "")
}

func GenerateRandomString(i int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, i)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
