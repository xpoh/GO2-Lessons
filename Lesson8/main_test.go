package main

import (
	"testing"
)

func TestGetFileList(t *testing.T) {
	list, err := getFileList("./")
	if err != nil {
		t.Error("err!=nil")
	}
	if len(list) != 3 {
		t.Error("len(list)!=3")
	}
	if list[0].fs.Name() != "main.go" {
		t.Error("list[0].fs.Name()!=main.go")
	}
	if list[1].fs.Name() != "main_test.go" {
		t.Error("list[1].fs.Name()!=main_test.go")
	}
	if list[2].fs.Name() != "test1.txt" {
		t.Error("list[2].fs.Name()!=test.txt")
	}
}
