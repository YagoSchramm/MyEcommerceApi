package entity

import "github.com/google/uuid"

type Image struct {
	ID   uuid.UUID
	Path string
	Url  string
}
