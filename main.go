package main

import (
	"demo/struct/bins"
	"demo/struct/file"
	"demo/struct/storage"
	"fmt"
)

func main() {
	db := file.JsonDb{
		FileName: "data",
	}
	storage := storage.Storage{
		DB: db,
	}
	list := bins.NewBinList(storage)
	newId := bins.GetNewId(list)
	bin, err := bins.NewBin(newId, "sam", true)
	if err != nil {
		return
	}
	fmt.Println(bin)
}
