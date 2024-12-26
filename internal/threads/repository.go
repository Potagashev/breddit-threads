package threads

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ThreadRepository struct {
	DB *pgx.Conn
}

func NewThreadRepository(db *pgx.Conn) *ThreadRepository {
	return &ThreadRepository{DB: db}
}

func (r *ThreadRepository) Create(thread *ThreadWriteDB) (uuid.UUID, error) {
	var threadId uuid.UUID
	query := `INSERT INTO threads (title, text, user_id) VALUES ($1, $2, $3) RETURNING id as threadId`
	err := r.DB.QueryRow(context.Background(), query, thread.Title, thread.Text, thread.UserId).Scan(&threadId)
	if err != nil {
		return threadId, err
	}
	return threadId, nil
}

func (r *ThreadRepository) GetByID(id uuid.UUID) (*Thread, error) {
	thread := &Thread{}
	query := `SELECT id, title, text, user_id FROM threads WHERE id = $1`
	err := r.DB.QueryRow(context.Background(), query, id).Scan(&thread.ID, &thread.Title, &thread.Text, &thread.UserId)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (r *ThreadRepository) GetMany() ([]*Thread, error) {
	query := `SELECT id, title, text FROM threads`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []*Thread
	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.Text); err != nil {
			return nil, err
		}
		threads = append(threads, &thread)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return threads, nil
}

func (r *ThreadRepository) Update(thread *Thread) error {
	query := `UPDATE threads SET title = $1, text = $2 WHERE id = $3`
	_, err := r.DB.Exec(context.Background(), query, thread.Title, thread.Text, thread.ID)
	return err
}

func (r *ThreadRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM threads WHERE id = $1`
	_, err := r.DB.Exec(context.Background(), query, id)
	return err
}
