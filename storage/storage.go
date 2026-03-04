package storage

import (
	"demo/struct/file"
	"fmt"
)

func SaveBin(bytes []byte) {
	file.WriteFile(bytes, "data.json")
}

func TakeBin() ([]byte, error) {
	bytes, err := file.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes, nil
}
