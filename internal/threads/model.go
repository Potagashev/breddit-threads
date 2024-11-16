package threads

import "github.com/google/uuid"

type ThreadWrite struct {
    Title string
    Text string
}

type ThreadWriteResponse struct {
	ID uuid.UUID `json:"id"`
}

type Thread struct {
	ID      uuid.UUID `json:"id"`
    Title   string    `json:"title"`
    Text string    `json:"content"`
}
