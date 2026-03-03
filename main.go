package main

import (
	"fmt"
	"demo/struct/bins"
)

func main() {
	list := bins.BinList{}
	newId := bins.GetNewId(list)
	bin, err := bins.NewBin(newId, "sam", true)
	if err != nil {
		return
	}
	fmt.Println(bin)
}
