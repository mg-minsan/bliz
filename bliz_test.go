package main

import (
	"io/fs"
	"io/ioutil"
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
	baseDir := "/tmp/bliz"
	os.Mkdir(baseDir, fs.FileMode(0755))
	defer os.RemoveAll(baseDir)
	file := "/" + getPartitionFileName("key")
	err := ioutil.WriteFile(baseDir+file, []byte(`{"key": "value"}`), 0755)
	if err != nil {
		t.Error(err)
	}

	t.Run("Should return value", func(t *testing.T) {
		bliz := NewBliz(baseDir)
		if bliz.Get("key") != "value" {
			t.Error("cloud not found the value by key")
		}
	})

	t.Run("Should return empty string if not found", func(t *testing.T) {
		bliz := NewBliz(baseDir)
		if bliz.Get("key1") != "" {
			t.Error("return value while it should not")
		}
	})
}

func TestRead(t *testing.T) {
	baseDir := "/tmp/bliz"
	defer os.RemoveAll(baseDir)

	t.Run("Should set the value", func(t *testing.T) {
		bliz := NewBliz(baseDir)
		bliz.Set("key", "value")
		if bliz.Get("key") != "value" {
			t.Error("value not found by the key name key")
		}
	})
}
