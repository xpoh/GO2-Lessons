package main

import (
	"flag"
	"fmt"
)

type Config struct {
	FileName   string
	PacketSize int64
	SliceStep  int64
	Count      int64
}

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

func PrintHelp() {
	fmt.Println("Slice file util. ver 0.0")
}

func myPanic(str string) {
	r := recover()
	if r != nil {
		fmt.Println("Recovered: ", str)
	}
}
