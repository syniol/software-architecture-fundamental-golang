package storage

import (
	"os"
	"path/filepath"
)

type unix struct {
	path string
}

func newStorageUnix(rootPath string) Partitioner {
	return &unix{
		path: rootPath,
	}
}

func (u *unix) SaveFile(name string, content []byte) error {
	err := os.WriteFile(
		filepath.Join(u.path, name),
		content,
		644,
	)
	if err != nil {
		return err
	}

	return nil
}
