package repositories

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
)

type NoteStorageFile struct {
	path string
}

func NewNoteStorageFile(path string) *NoteStorageFile {
	return &NoteStorageFile{
		path: path,
	}
}

func (s *NoteStorageFile) Put(_ context.Context, key string, content *bytes.Buffer) error {
	// TODO: return uri to download file
	filename := path.Join(s.path, key)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("file %s was not created: %s", key, err)
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = content.WriteTo(file)
	if err != nil {
		return fmt.Errorf("can not write content to the file %s: %s", key, err)
	}

	return nil
}
