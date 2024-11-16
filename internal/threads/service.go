package threads

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrThreadNotFound = errors.New("thread not found")
)

type ThreadService struct {
	ThreadRepository *ThreadRepository
}

func NewThreadService(threadRepo *ThreadRepository) *ThreadService {
	return &ThreadService{ThreadRepository: threadRepo}
}

func (s *ThreadService) CreateThread(thread *ThreadWrite) (uuid.UUID, error) {
	return s.ThreadRepository.Create(thread)
}

func (s *ThreadService) GetThreadByID(id uuid.UUID) (*Thread, error) {
	thread, err := s.ThreadRepository.GetByID(id)
	if err != nil {
		return nil, ErrThreadNotFound
	}
	return thread, nil
}

func (s *ThreadService) GetManyThreads() ([]*Thread, error) {
	threads, err := s.ThreadRepository.GetMany()
	if err != nil {
		return nil, err
	}

	if threads == nil {
		threads = []*Thread{}
	}
	return threads, nil
}

func (s *ThreadService) UpdateThread(thread *Thread) error {
	if _, err := s.ThreadRepository.GetByID(thread.ID); err != nil {
		return ErrThreadNotFound
	}
	return s.ThreadRepository.Update(thread)
}

func (s *ThreadService) DeleteThread(id uuid.UUID) error {
	if _, err := s.ThreadRepository.GetByID(id); err != nil {
		return ErrThreadNotFound
	}
	return s.ThreadRepository.Delete(id)
}
