package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
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

func printHelp() {
	fmt.Println("Программа удаления дублированных файлов.")
	fmt.Println("аргументы:.")
	fmt.Println("-h 	текщая справка")
	fmt.Println("-p 	путь для поиска файлов")
	fmt.Println("-f 	удаление с подтверждением повторяющихся файлов")
}

func (f *fileInfoWithPath) removeFile(wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	err = os.Remove(f.path)
	if err == nil {
		fmt.Println("Удален файл: ", f.path)
	}
	return err
}

func main() {
	wg := sync.WaitGroup{}
	fConfirm := false
	dir := flag.String("p", "", "путь для поиска файлов")
	fDel := flag.Bool("f", false, "удаление с подтверждением повторяющихся файлов")
	fHelp := flag.Bool("h", false, "текщая справка")
	flag.Parse()

	if *fHelp {
		printHelp()
		return
	}

	if *fDel {
		fmt.Print("Удалять файлы сразу? [напиши YES] ")
		var s string
		_, _ = fmt.Scanln(&s)

		if s == "YES" {
			fConfirm = true
		}
	}

	list, err := getFileList(*dir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Найдено файлов: ", len(list))
	m := make(map[string]string)

	for i := 0; i < len(list); i++ {
		idx := list[i].fs.Name() + "_" + strconv.FormatInt(list[i].fs.Size(), 10)
		_, ok := m[idx]
		if !ok {
			m[idx] = list[i].path
			fmt.Println("Uniq file: ", list[i].path)
		} else if fConfirm {
			wg.Add(1)
			go list[i].removeFile(&wg)
		}
	}
	wg.Wait()
}
