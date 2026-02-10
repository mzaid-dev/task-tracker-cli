package storage

import (
	"errors"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}

func CreateFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	return file.Close()
}

func WriteFile(filename string, content string) error {

	return os.WriteFile(filename, []byte(content), 0644)

}

func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(content), nil

}

func DeleteFile(filename string) error {
	return os.Remove(filename)
}
