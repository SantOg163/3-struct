package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}
type BinList = []Bin

func getNewId(list BinList) string {
	if len(list) >= 1 {
		lastElementId := (list[len(list)-1].id)

		lastElementIntId, _ := strconv.Atoi(lastElementId)

		return strconv.Itoa(lastElementIntId + 1)

	} else {
		return "1"
	}
}
func newBin(id string, name string, private bool) (*Bin, error) {
	if id == "" || name == ""{
		return nil, errors.New("Empty id or name")
	}
	currentBin := Bin{
		id:      id,
		name:    name,
		private: private,
		createdAt: time.Now(),
	}
	return &currentBin,nil

}
func main() {
	list := BinList{}
	newId := getNewId(list)
	bin, err := newBin(newId, "sam", true)
	if err != nil {
		return
	}
	fmt.Println(bin)
}
