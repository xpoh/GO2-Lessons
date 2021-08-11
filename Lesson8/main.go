package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
)

type fileInfoWithPath struct {
	fs   fs.FileInfo
	path string
}

func getFileList(path string) (list []fileInfoWithPath, err error) {
	var result []fileInfoWithPath
	var tmp = fileInfoWithPath{
		fs:   nil,
		path: "",
	}

	fileStat, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}
	for i := 0; i < len(fileStat); i++ {
		if fileStat[i].IsDir() {
			tmp, err := getFileList(path + "/" + fileStat[i].Name())
			if err == nil {
				for j := 0; j < len(tmp); j++ {
					result = append(result, tmp[j])
				}
			} else {
				return nil, err
			}
		} else {
			tmp.path = path + "/" + fileStat[i].Name()
			tmp.fs = fileStat[i]
			result = append(result, tmp)
		}
	}
	return result, nil
}

func main() {
	dir := flag.String("path", "", "directory for search")
	flag.Parse()
	list, err := getFileList(*dir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Найдено файлов: ", len(list))
	m := make(map[string]string)
	forDel := make(map[string]string)
	for i := 0; i < len(list); i++ {
		idx := list[i].fs.Name() + "_" + strconv.FormatInt(list[i].fs.Size(), 10)
		_, ok := m[idx]
		if !ok {
			m[idx] = list[i].path
		} else {
			forDel[idx] = list[i].path
		}
	}
	fmt.Println("Uniq files count:", len(m))
	fmt.Println("For delete files count:", len(forDel))
	fmt.Print("Удалить ", len(forDel), " файлов? [напиши YES] ")
	var s string
	fmt.Scanln(&s)

	if s == "YES" {
		count := 0
		fmt.Println("Удаление файлов:")
		for _, p := range forDel {
			fmt.Print("Удаление ", p)
			if os.Remove(p) == nil {
				fmt.Println(" [Ok]")
				count++
			} else {
				fmt.Println(" [Error]")
			}
		}
		println("Удалено ", count, "файлов")
	}
}
