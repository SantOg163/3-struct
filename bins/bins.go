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
type BinList []Bin

func GetNewId(list *BinList) string {
	if len(*list) >= 1 {
		lastElementId := ((*list)[len(*list)-1].Id)

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
func NewBinList() *BinList {
	bytes, err := storage.TakeBin()
	if err != nil {
		return &BinList{}
	}
	result := BinList{}
	err = json.Unmarshal(bytes, &result)
	return &result
}

func (list *BinList) AddBin(bin *Bin) {
	*list = append((*list), *bin)

	bytes, err := json.Marshal(list)
	if err != nil {
		fmt.Println("Произошла ошибка при переводе в байты")
		return
	}
	storage.SaveBin(bytes)
}
