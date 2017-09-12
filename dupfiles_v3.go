package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var p sync.Map

func fileHashToHexString(file *os.File) string {
	p1 := md5.New()
	if _, err := io.Copy(p1, file); err != nil {
		log.Fatal(err)
	}

	q := hex.EncodeToString(p1.Sum(nil))
	return q
}
func compareAndReturn(input []string) {
	//solution := make(map[string][]string)
	for i := 0; i < len(input); i++ {
		f1, _ := os.Open(input[i])
		r := fileHashToHexString(f1)
		if val, ok := p.Load(r); !ok {
			p.Store(r, []string{input[i]})
		} else {
			ip, ok := val.([]string)
			if ok {
				ip = append(ip, input[i])
				p.Store(r, ip)
			}
		}

	}

}
func f1(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	path = strings.Replace(path, "\n", "", -1)
	if path == "" {
		path = "./"
	}

	listFiles(path, wg)

}

func listFiles(path string, wg *sync.WaitGroup) {
	fileNames, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	files := []string{}
	directories := []string{}

	for _, f := range fileNames {
		if !f.IsDir() {
			files = append(files, filepath.Join(path, f.Name()))
		} else {
			directories = append(directories, filepath.Join(path, f.Name()))
		}
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		compareAndReturn(files)
	}()

	for _, dirs := range directoriess {
		listFiles(dirs, wg)
	}
}
func main() {
	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	text := "/Users/akshaydeo/Downloads"
	f1(text, &wg)
	wg.Wait()
	s := make(map[string][]string)
	fmt.Println("The Duplicate Files are:\n")
	p.Range(func(k string, v []string) bool {
		s[k] = v
		if len(v > 1) {
			fmt.Println(">>>values:\n", v)
		}

		return true
	})

}
