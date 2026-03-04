package main

import (
	"demo/struct/bins"
	"demo/struct/file"
	"fmt"
)

func main() {
	res, _ := file.ReadFile("test.txt")
	fmt.Println(res)
	list := bins.NewBinList()
	newId := bins.GetNewId(list)
	bin, err := bins.NewBin(newId, "sam", true)
	if err != nil {
		return
	}
	fmt.Println(bin)
}
