package main

import (
	"demo/struct/api"
	"demo/struct/bins"
	"demo/struct/config"
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
	conf := config.NewConfig()
	api := api.NewApi(*conf)
	fmt.Println(api)
	fmt.Println(bin)
}
