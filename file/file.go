package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonDb struct {
	FileName string
}

func (db JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return nil, err
	}
	return (data), nil
}
func (db JsonDb) Write(content []byte) error {
	if !strings.Contains(db.FileName, ".json") {
		fmt.Println("Доступен только json")
		return errors.New("Доступен только json")
	}
	file, err := os.Create(db.FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	fmt.Println("Запись в файл успешна")
	return nil
}
