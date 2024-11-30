package storage

import (
	"os"
	"path/filepath"
)

type Unix struct{}

func (*Unix) SaveFile(name, path string, content []byte) error {
	err := os.WriteFile(
		filepath.Join(path, name),
		content,
		644,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewStorageUnix() Partitioner {
	return &Unix{}
}
