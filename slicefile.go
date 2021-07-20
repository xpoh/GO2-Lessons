// Программа для нарезки файла на части методом откидывания начального участка с данными
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// SliceError стуктура для хранения расширенной версии ошибки
type SliceError struct {
	err       error
	indexFile int64
	timeError time.Time
	msg       string
}

func (e *SliceError) Error() string {
	return fmt.Sprintf(e.msg, e.err, e.timeError, e.indexFile)
}

func (e *SliceError) Unwrap() error {
	return e.err
}

// WrapSliceError функция обертка, дополняющая сообщение об ошибке дополнитльной информацией
func WrapSliceError(err error, indexFile int64, timeError time.Time, msg string) error {
	return &SliceError{
		err:       err,
		indexFile: indexFile,
		timeError: timeError,
		msg:       msg,
	}
}

// SliceFile функция осуществляющая нарезку файла в соответствии с переданной конфигурацией
func SliceFile(c Config) error {
	data, err := ioutil.ReadFile(c.FileName)
	if err != nil {
		return err
	}

	for i := int64(0); i < c.Count; i++ {
		var destName string = c.FileName + "-" + strconv.FormatInt(i+1, 10) + ".dat"
		fmt.Println("Creating...", destName)

		dest, err := os.Create(destName)
		if err != nil {
			return err
		}

		n, err := dest.Write(data[(c.PacketSize * c.SliceStep * i):len(data)])
		if err != nil {
			return WrapSliceError(err, i, time.Now(), "Error Write function!!!")
		}
		fmt.Println("written ", n, " bytes")

		err = dest.Close()
		if err != nil {
			return WrapSliceError(err, i, time.Now(), "Error Write function!!!")
		}
	}
	return nil
}
