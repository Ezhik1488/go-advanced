package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type LocalStorage struct {
	filename string
	Data     map[string]string
}

func NewLocalStorage(filename string) *LocalStorage {
	ls := &LocalStorage{filename: filename}

	rawData, err := ls.Read()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(rawData, &ls.Data)
	if err != nil {
		ls.Data = make(map[string]string)
	}

	return ls
}

func (s *LocalStorage) Write(rawData map[string]string) error {

	for key, value := range rawData {
		if _, ok := s.Data[key]; ok {
			return errors.New(" email already exists")
		}
		s.Data[key] = value
	}

	data, err := json.MarshalIndent(s.Data, "", "\t")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(s.filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
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

func (s *LocalStorage) Save() error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	data, err := json.MarshalIndent(s.Data, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
