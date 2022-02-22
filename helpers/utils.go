package helpers

import (
	"crypto/rand"
	"fmt"
)

func Create_unique_string(length int) string {
	bytestring := make([]byte, length)
	if _, err := rand.Read(bytestring); err != nil {
		panic(err)
	}
	// todo: query db to check if string exists already or not
	str := fmt.Sprintf("%X", bytestring)
	return str
}