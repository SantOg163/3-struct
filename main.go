package main

import (
	"demo/struct/api"
	"demo/struct/bins"
	"demo/struct/config"
	"demo/struct/storage"
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Ошибка загрузки .env")
	}

	isCreate := flag.Bool("create", false, "isCreate")
	isUpdate := flag.Bool("update", false, "isUpdate")
	isDelete := flag.Bool("delete", false, "isDelete")
	isGet := flag.Bool("get", false, "isGer")
	isList := flag.Bool("list", false, "isList")
	fileName := flag.String("file", "", "isList")
	id := flag.String("id", "", "id")
	flag.Parse()
	if (*isCreate || *isUpdate) && *fileName == "" {
		panic("Нет данных для отправки")
	}
	if (*isUpdate || *isDelete || *isGet) && *id == "" {
		panic("Нет id")

	}

	myConfig := config.NewConfig()
	myApi := api.NewApi(*myConfig)
	storage := storage.NewStorage("data", *myApi)
	list := bins.NewBinList(storage)
	if *isCreate {
		myApi.CreateBin(*fileName)
	}
	if *isUpdate {
		myApi.UpdateBin(*fileName, *id)
	}
	if *isDelete {
		myApi.DeleteBin(*id)
	}
	if *isGet {
		myApi.GetBin(*id)
	}
	if *isList {
		fmt.Println(list)
	}

	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Ошибка загрузки .env")
	// }
	// conf := config.NewConfig()
	// myApi := api.NewApi(*conf)
	// stor := storage.NewStorage("data.json", *myApi)
	// list := bins.NewBinList(stor)
	// newId := bins.GetNewId(list)
	// bin, err := bins.NewBin(newId, "sam", true)
	// if err != nil {
	// 	return
	// }
	// list.AddBin(bin)
}
