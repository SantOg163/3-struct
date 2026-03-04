package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return (data), nil
}
func WriteFile(content []byte, name string) error {
	if !strings.Contains(name,".json"){
		fmt.Println("Доступен только json")
		return errors.New("Доступен только json")
	}
	file, err := os.Create(name)
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
