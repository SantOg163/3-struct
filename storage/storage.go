package storage

import (
	"demo/struct/file"
	"fmt"
)

type Storage struct {
	DB file.JsonDb
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
