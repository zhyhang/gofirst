package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please input lines, CTRL+D will end input.")
	counter := make(map[string]int)
	//	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	// for line, err := reader.ReadString('\n'); err == nil && len(line) > 1; line, err = reader.ReadString('\n') {
	for scanner.Scan() {
		counter[scanner.Text()]++
		//		fmt.Println(scanner.Text())
	}
	for line, n := range counter {
		fmt.Printf("%s-->%d\n", line, n)
	}
}
