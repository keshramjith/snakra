package id_generation

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/hailongz/golang/basex"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func NewId() (uuid.UUID, string) {

	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
	}
	return uuid, convertUUIDToBase63(uuid)
}

func convertUUIDToBase63(id uuid.UUID) string {
	var base62, _ = basex.NewEncoding(alphabet)
	return base62.Encode(id.Bytes())
}
