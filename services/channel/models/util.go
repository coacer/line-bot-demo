package models

import "github.com/google/uuid"

func GeneratePrimaryKey() string {
	uuidObj, _ := uuid.NewUUID()
	return uuidObj.String()
}
