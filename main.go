package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataPath := fmt.Sprintf("%s/.config/bliz/data", homeDir)

	dat, err := ioutil.ReadFile(dataPath)
	if err != nil {
		path := fmt.Sprintf("%s/.config/bliz", homeDir)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.Mkdir(path, fs.FileMode(0755))
			if err != nil {
				panic(err)
			}
		}

		dat = []byte("{}")
		err := ioutil.WriteFile(dataPath, []byte("{}"), 0755)
		if err != nil {
			panic(err)
		}
	}

	hash := make(map[string]string)
	err = json.Unmarshal(dat, &hash)
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "get":
		checkForKey(os.Args)
		key := os.Args[2]
		if checkForKey(os.Args) {
			return
		}
		if val, ok := hash[key]; ok {
			fmt.Println(val)
			return
		}
		fmt.Println("key not found")
		return
	case "set":
		if checkForKey(os.Args) {
			return
		}
		key := os.Args[2]
		if key == "" {
			fmt.Println("key not found")
			return
		}
		if len(os.Args) < 4 {
			fmt.Println("value is not set")
			return
		}
		value := os.Args[3]
		hash[key] = value
		writeToFile(hash, dataPath)
	case "list":
		for key := range hash {
			fmt.Println(key)
		}
	default:
		fmt.Println("command not found")
	}
}

func writeToFile(data map[string]string, filePath string) {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON parsed error, file is corrected")
		os.Exit(1)
	}
	ioutil.WriteFile(filePath, json, 0644)
}

func errorKeyRequried() {
	fmt.Println("key is required")
}

func checkForKey(args []string) bool {
	if len(os.Args) < 3 {
		errorKeyRequried()
		return true
	}
	return false
}
