package storage

import (
	"demo/struct/api"
	"demo/struct/file"
	"fmt"
)

type Storage struct {
	DB file.JsonDb
}

func NewStorage(name string, api api.Api) *Storage {
	return &Storage{
		DB: file.JsonDb{
			FileName: name,
		},
	}
}

func (stor *Storage) Write(bytes []byte) {
	stor.DB.Write(bytes)
}

func (stor *Storage) Read() ([]byte, error) {
	bytes, err := stor.DB.Read()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes, nil
}
