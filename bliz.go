package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Bliz interface {
	Get()
	Set()
	List()
}

type BlizJson struct {
	dataFilePath string
	data         map[string]string
}

// Get return value by key
func (bliz BlizJson) Get(key string) string {
	return bliz.data[key]
}

func (bliz *BlizJson) Set(key string, value string) {
	log.Print(bliz.data)
	bliz.data[key] = value
	if err := writeToFile(bliz.data, bliz.dataFilePath); err != nil {
		fmt.Print(err.Error())
	}
}

func (bliz BlizJson) List() []string {
	keys := make([]string, len(bliz.data))
	for key := range bliz.data {
		keys = append(keys, key)
	}
	return keys
}

func NewBliz(filePath string) *BlizJson {
	if _, err := os.Stat(filePath); os.IsExist(err) {
		createFileRecursively(filePath)
	}
	data, err := ioutil.ReadFile(filePath)
	hash := make(map[string]string)
	err = json.Unmarshal(data, &hash)

	if err != nil {
		panic(err)
	}

	return &BlizJson{
		dataFilePath: filePath,
		data:         hash,
	}
}

func createFileRecursively(filePath string) {
	fileDir := strings.Split(filePath, "/")
	for i := range fileDir {

		targetFile := strings.Join(fileDir[:i+1], "/")
		if targetFile == "" {
			continue
		}
		if _, err := os.Stat(targetFile); os.IsNotExist(err) {
			if i == len(fileDir)-1 {
				err = ioutil.WriteFile(targetFile, []byte("{}"), 0755)
				if err != nil {
					panic(err)
				}
				continue
			}

			err = os.Mkdir(targetFile, fs.FileMode(0755))
			if err != nil {
				panic(err)
			}
		}
	}
}

func writeToFile(data map[string]string, filePath string) error {
	log.Println(data)
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filePath, json, 0755)
	return nil
}
