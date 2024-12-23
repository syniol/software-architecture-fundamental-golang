package storage

import (
	"os"
	"path/filepath"
)

type unix struct {
	path string
}

func newStorageUnix(rootPath string) Manager {
	return &unix{
		path: rootPath,
	}
}

func (u *unix) SaveFile(name string, content []byte) error {
	err := os.WriteFile(
		filepath.Join(u.path, name),
		content,
		0777,
	)
	if err != nil {
		return err
	}

	return nil
}
