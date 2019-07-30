package main

import (
	"io/ioutil"
	"strings"
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("expected clear-file-separator src.file [dst.file]")
	}
	if len(os.Args) > 2 {
		err := clear(os.Args[1], os.Args[2])
		if err != nil {
			panic(err)
		}
	} else {
		files, err := ioutil.ReadDir(os.Args[1])
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Printf("file: %s\n", file.Name())
			err = clear(os.Args[1]+file.Name(), os.Args[1]+file.Name())
			if err != nil {
				panic(err)
			}
		}
	}
}

func clear(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("is dir")
	}
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	s := strings.Replace(string(data), string(rune(13)), "", -1)
	return ioutil.WriteFile(dst, []byte(s), info.Mode())
}
