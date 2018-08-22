package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	filePath := "/tmp/mmap.test.dat"
	// open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// write data to file and flush, lead to mmap byte[] is available read
	writer := bufio.NewWriter(file)
	writer.WriteByte(1)
	writer.WriteByte(2)
	writer.Flush()
	maxFileSize := 1024 * os.Getpagesize()
	// open a mmap file
	fileData, err1 := syscall.Mmap(int(file.Fd()), 0, maxFileSize, syscall.PROT_READ, syscall.MAP_SHARED) //syscall.MAP_SHARED)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("open mmap file " + filePath + " OK, len: " + strconv.Itoa(len(fileData)) +
		", first byte: " + strconv.Itoa(int(fileData[0])))
	syscall.Munmap(fileData)

}
