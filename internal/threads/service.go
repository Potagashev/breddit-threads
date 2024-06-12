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

// NewThreadService создает новый экземпляр ThreadService
func NewThreadService(threadRepo *ThreadRepository) *ThreadService {
    return &ThreadService{ThreadRepository: threadRepo}
}

// CreateThread создает новый тред
func (s *ThreadService) CreateThread(thread *Thread) error {
    return s.ThreadRepository.Create(thread)
}

// GetThreadByID получает тред по ID
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
    return threads, nil
}


// UpdateThread обновляет данные треда
func (s *ThreadService) UpdateThread(thread *Thread) error {
    if _, err := s.ThreadRepository.GetByID(thread.ID); err != nil {
        return ErrThreadNotFound
    }
    return s.ThreadRepository.Update(thread)
}

// DeleteThread удаляет тред по ID
func (s *ThreadService) DeleteThread(id uuid.UUID) error {
    if _, err := s.ThreadRepository.GetByID(id); err != nil {
        return ErrThreadNotFound
    }
    return s.ThreadRepository.Delete(id)
}
