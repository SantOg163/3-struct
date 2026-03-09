package api

import (
	"bytes"
	"demo/struct/bins"
	"demo/struct/config"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Api struct {
	conf config.Config
}

func NewApi(conf config.Config) *Api {
	return &Api{conf: conf}
}

func (api *Api) CreateBin(fileName string) ([]byte, error) {
	var newBin bins.Bin
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	json.Unmarshal(file, &newBin)
	data, err := json.Marshal(newBin)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", api.conf.MasterKey)
	req.Header.Add("X-Access-Key", api.conf.AccessKey)
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println("Not created: ",resp.StatusCode)
		return nil, errors.New("Not200")
	}

	newData, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return newData, nil

}
func (api *Api) DeleteBin(binId string) (bool, error) {
	req, err := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+binId, nil)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", api.conf.MasterKey)
	req.Header.Add("X-Access-Key", api.conf.AccessKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer resp.Body.Close()

	result := resp.StatusCode == 200 || resp.StatusCode == 201

	return result, nil
}
func (api *Api) GetBin(binId string) ([]byte, error) {

	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+binId, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", api.conf.MasterKey)
	req.Header.Add("X-Access-Key", api.conf.AccessKey)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
func (api *Api) UpdateBin(fileName string, binId string) ([]byte, error) {
	var newData bins.Bin
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	json.Unmarshal(file, &newData)

	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+binId, bytes.NewBuffer(file))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", api.conf.MasterKey)
	req.Header.Add("X-Access-Key", api.conf.AccessKey)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil

}
