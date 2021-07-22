package slice

import (
	"flag"
	"fmt"
)

// Config структура с основными параметрами нарезки файла:
//	FileNme - Имя файла для нарезки
//	PacketSize размер пакета данных
//	SliceStep  шаг в количествах пакетов с которым осуществляется нарезка
//	Count количество файлов
type Config struct {
	FileName   string
	PacketSize int64
	SliceStep  int64
	Count      int64
}

// ReadConfig возвращает структуру данных с конфигурацией работы программы
func ReadConfig() (*Config, error) {
	defer myPanic("Read config error (mypanic)")

	var fileName = flag.String("if", "IMU.dat", "file for slicing")
	var packetSize = flag.Int64("bsize", 44, "packet size for slicing")
	var sliceStep = flag.Int64("step", 0, "sliceDat interval")
	var err = error(nil)
	var count = flag.Int64("count", 100, "sliceDat count")

	flag.Parse() // В исходном коде есть вариант паники, но как вызвать его я не придумал
	//panic("TEST PANIC")
	c := Config{
		FileName:   *fileName,
		PacketSize: *packetSize,
		SliceStep:  *sliceStep,
		Count:      *count,
	}
	return &c, err
}

// PrintHelp выводит в стандартный поток вывода справку по работе с программой
func PrintHelp() {
	fmt.Println("Slice file util. ver 0.0")
}

func myPanic(str string) {
	r := recover()
	if r != nil {
		PrintHelp()
		fmt.Println("Recovered: ", str)
	}
}
