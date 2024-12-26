package threads

import "github.com/google/uuid"

type ThreadWrite struct {
    Title string
    Text string
}

type ThreadWriteDB struct {
    ThreadWrite
    UserId uuid.UUID
}

type ThreadWriteResponse struct {
	ID uuid.UUID `json:"id"`
}

type Thread struct {
	ID      uuid.UUID `json:"id"`
    UserId  uuid.UUID `json:"userId"`
    Title   string    `json:"title"`
    Text string    `json:"content"`
}
