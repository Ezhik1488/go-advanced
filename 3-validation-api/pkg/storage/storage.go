package storage

import (
	"encoding/json"
	"os"
)

type LocalStorage struct {
	filename string
}

func NewLocalStorage(filename string) *LocalStorage {
	return &LocalStorage{filename: filename}
}

func (s *LocalStorage) Write(rawData map[string]string) error {
	data, err := json.MarshalIndent(rawData, "", "\t")
	if err != nil {
		return err
	}
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *LocalStorage) Read() ([]byte, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
