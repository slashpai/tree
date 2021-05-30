package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	fmt.Println(args)

	for _, arg := range args {
		err := tree(arg)

		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}

}

func tree(path string) error {
	fmt.Println(path)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		fmt.Printf("%q", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %v\n", err)
		return err
	}

	return nil
}
