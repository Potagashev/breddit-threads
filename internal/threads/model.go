package threads

import "github.com/google/uuid"

type Thread struct {
	ID    uuid.UUID
	Title string
	Text  string
}