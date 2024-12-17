package storage

import (
	"log"
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
	log.Println("inside save file method")
	log.Println("file path", u.path)
	log.Println("file name", name)

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
