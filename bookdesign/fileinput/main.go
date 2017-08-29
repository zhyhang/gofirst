package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("not specify the input files, using console input! linux: CTRL+D end input windows: CTRL+Z")
		countLines(os.Stdin, counter)
	} else {
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("open file %s error: %v.", file, err)
				continue
			}
			countLines(f, counter)
			f.Close()
		}
	}
	for line, n := range counter {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counter map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counter[scanner.Text()]++
	}
}
