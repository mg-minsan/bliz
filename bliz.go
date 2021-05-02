package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Bliz interface {
	Get()
	Set()
	List()
}

type BlizJson struct {
	dataFilePath string
	filepath     string
}

// Get return value by key
func (bliz *BlizJson) Get(key string) string {
	bliz.setPartitionedFilePath(key)
	if exist := pathExist(bliz.filepath); !exist {
		return ""
	}
	data := parsedFile(bliz.filepath)
	return data[key]
}

func (bliz *BlizJson) Set(key string, value string) {
	bliz.setPartitionedFilePath(key)
	if exist := pathExist(bliz.filepath); !exist {
		createFileRecursively(bliz.filepath)
	}
	data := parsedFile(bliz.filepath)
	data[key] = value
	if err := writeToFile(data, bliz.filepath); err != nil {
		fmt.Print(err.Error())
	}
}

func (bliz *BlizJson) setPartitionedFilePath(key string) {
	bliz.filepath = bliz.dataFilePath + "/" +
		getPartitionFileName(key)
}

func NewBliz(filePath string) *BlizJson {
	if _, err := os.Stat(filePath); os.IsExist(err) {
		createFileRecursively(filePath)
	}
	//data, err := ioutil.ReadFile(filePath)
	/* hash := make(map[string]string)*/
	//err = json.Unmarshal(data, &hash)

	//if err != nil {
	//panic(err)
	/* }*/

	return &BlizJson{
		dataFilePath: filePath,
	}
}

func parsedFile(filePath string) map[string]string {
	data, err := ioutil.ReadFile(filePath)
	hash := make(map[string]string)
	err = json.Unmarshal(data, &hash)
	if err != nil {
		panic(err)
	}
	return hash
}

func createFileRecursively(filePath string) {
	fileDir := strings.Split(filePath, "/")
	for i := range fileDir {

		targetFile := strings.Join(fileDir[:i+1], "/")
		if targetFile == "" {
			continue
		}
		if _, err := os.Stat(targetFile); os.IsNotExist(err) {
			if filepath.Ext(targetFile) == ".json" {
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

func pathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeToFile(data map[string]string, filePath string) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filePath, json, 0755)
	return nil
}

func getPartition(key rune) int {
	return int(key) % 10
}

func getPartitionFileName(key string) string {
	partitionkey := getPartition(rune(key[0]))
	return strconv.Itoa(partitionkey) + ".json"
}
