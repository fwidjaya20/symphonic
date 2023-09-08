package file

import (
	"os"
	"path/filepath"
)

func Create(file string, content string) error {
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}

func Exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
