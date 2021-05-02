package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateFileIfNotExit(t *testing.T) {
	filePath := "/tmp/bliz/data"
	defer os.RemoveAll("/tmp/bliz")
	createFileRecursively(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("file was not created")
	}
}

func TestGet(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "bliz_data")
	defer os.Remove(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(file.Name(), []byte(`{"key": "value"}`), 0755)

	t.Run("Should return value", func(t *testing.T) {
		bliz := NewBliz(file.Name())
		if bliz.Get("key") != "value" {
			t.Error("cloud not found the value by key")
		}
	})

	t.Run("Should return empty string if not found", func(t *testing.T) {
		bliz := NewBliz(file.Name())
		if bliz.Get("key1") != "" {
			t.Error("return value while it should not")
		}
	})

}

func TestList(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "bliz_data")
	defer os.Remove(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(file.Name(), []byte(`{"key": "value"}`), 0755)
	t.Run("Should return array of keys", func(t *testing.T) {
		bliz := NewBliz(file.Name())
		list := bliz.List()
		if list[1] != "key" {
			t.Error("should return value from ", list)
		}
	})
}

func TestRead(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "bliz_data")
	defer os.Remove(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(file.Name(), []byte(`{}`), 0755)
	t.Run("Should set the value", func(t *testing.T) {
		bliz := NewBliz(file.Name())
		bliz.Set("key", "value")
		if bliz.data["key"] != "value" {
			t.Error("value not found by the key name key")
		}
	})
}
