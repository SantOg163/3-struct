package bins

import (
	"demo/struct/storage"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
type BinList struct {
	Bins    []Bin
	storage storage.Storage
}
type Data interface{
	Write(bytes []byte)
	Read() ([]byte, error)
}

func GetNewId(list *BinList) string {
	if len(list.Bins) >= 1 {
		lastElementId := list.Bins[len(list.Bins)-1].Id

		lastElementIntId, _ := strconv.Atoi(lastElementId)

		return strconv.Itoa(lastElementIntId + 1)

	} else {
		return "1"
	}
}

func NewBin(id string, name string, private bool) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("Empty id or name")
	}
	currentBin := &Bin{
		Id:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
	field, _ := reflect.TypeOf(currentBin).Elem().FieldByName("name")
	fmt.Println(string(field.Tag))

	return currentBin, nil

}
func NewBinList(stor storage.Storage) *BinList {
	bytes, err := stor.Read()
	if err != nil {
		return &BinList{
			Bins:    []Bin{},
			storage: stor,
		}
	}
	result := BinList{}
	err = json.Unmarshal(bytes, &result)
	result.storage = stor
	return &result
}

func (list *BinList) AddBin(bin *Bin) {
	list.Bins = append(list.Bins, *bin)

	bytes, err := json.Marshal(list.Bins)
	if err != nil {
		fmt.Println("Произошла ошибка при переводе в байты")
		return
	}
	list.storage.Write(bytes)
}
