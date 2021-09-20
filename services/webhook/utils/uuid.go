package utils

import "github.com/google/uuid"

func GenerateUuid() string {
	uuidObj, _ := uuid.NewUUID()
	return uuidObj.String()
}
