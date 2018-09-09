package main

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"os"
)

// console command like bash tail
func main() {
	follow1(os.Args[1])
}

func follow1(filename string) error {
	file, _ := os.Open(filename)
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	watcher.Add(filename)
	r := bufio.NewReader(file)
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Print(string(by))
		if err != io.EOF {
			continue
		}
		if err = waitForChange1(watcher); err != nil {
			return err
		}
	}
}

func waitForChange1(w *fsnotify.Watcher) error {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-w.Errors:
			return err
		}
	}
}
