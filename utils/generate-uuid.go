package utils

import (guuid "github.com/google/uuid" )

func GenerateUuid() string {
	return guuid.New().String()
}