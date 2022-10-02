package storage

import (
	"errors"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
	"sync"
)

type Data map[string]*domain.Event

type Storage struct {
	sync.RWMutex
	Data
}

func (s *Storage) Set(id string, data domain.Event) {
	s.Lock()
	s.Data[id] = &data
	s.Unlock()
}

func (s *Storage) Get(id string) (*domain.Event, error) {
	s.RLock()
	defer s.RUnlock()
	if data, ok := s.Data[id]; ok {
		return data, nil
	}
	return nil, errors.New("event not found")
}

func (s *Storage) Update(id string, calendar domain.Event) error {
	data, err := s.Get(id)
	if err != nil {
		return err
	}

	s.Lock()
	if calendar.Id != "" {
		data.Id = calendar.Id
	}

	if calendar.Title != "" {
		data.Title = calendar.Title
	}

	if calendar.Description != "" {
		data.Description = calendar.Description
	}

	if !calendar.Date.IsZero() {
		data.Date = calendar.Date
	}

	if calendar.UserId != 0 {
		data.UserId = calendar.UserId
	}

	s.Unlock()
	return nil

}

func (s *Storage) Delete(id string) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	s.Lock()
	delete(s.Data, id)
	s.Unlock()
	return nil
}
