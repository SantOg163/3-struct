package main

import (
	"demo/struct/bins"
	"demo/struct/storage"
	"fmt"
)

func main() {
	storage := storage.NewStorage("data")
	list := bins.NewBinList(storage)
	newId := bins.GetNewId(list)
	bin, err := bins.NewBin(newId, "sam", true)
	if err != nil {
		return
	}
	fmt.Println(bin)
}
