package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetFileName(path string) ([]string, error) {
	var names []string
	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		names = append(names, file.Name())
		return err
	})
	return names, err
}

func main() {
	names, err := GetFileName("./")
	if err != nil {
		log.Fatal(err)
		return
	}

	for num, name := range names {
		fmt.Printf("%v:%v\n", num, name)
	}
	return
}
