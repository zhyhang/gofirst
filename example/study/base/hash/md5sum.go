package hash

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

func Md5Sum() {
	data := []byte("These pretzels are making me thirsty.")
	md5Bytes := md5.Sum(data)
	fmt.Printf("%x\n", md5Bytes)
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%x", md5Bytes)
	fmt.Println(buf.String())
	md5s := fmt.Sprintf("%x", md5Bytes)
	fmt.Println(md5s)
}
