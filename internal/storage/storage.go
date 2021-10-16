package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kosenkovd/storage/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	if foundFile, ok := s.files[fileID]; ok {
		return foundFile, nil
	}

	return nil, fmt.Errorf("File %v not found.", fileID)
}
