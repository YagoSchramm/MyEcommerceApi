package util

import "github.com/google/uuid"

func GenerateID() *uuid.UUID {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}
	return &id
}
