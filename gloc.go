package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type result struct {
	filename  string
	total     int64
	empty     int64
	effective int64
	comment   int64
}

func walkDir(dir string, files chan<- string) {
	for _, entry := range dirents(dir) {
		absolute := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			walkDir(absolute, files)
		} else {
			files <- absolute
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gloc: %v\n", err)
	}
	return entries
}

var workers = map[string]func(string) result{
	".c":   cxx,
	".cc":  cxx,
	".cpp": cxx,
	".h":   cxx,
	".hpp": cxx,
}

var project string

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// traverse the file tree
	files := make(chan string)
	go func() {
		for _, root := range roots {
			project = root
			walkDir(root, files)
		}
		close(files)
	}()

	// calculate loc
	retChan := make(chan result)
	waitGroup := sync.WaitGroup{}
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		waitGroup.Add(1);
		go func() {
			for f := range files {
				extension := filepath.Ext(f)
				if len(extension) != 0 && workers[extension] != nil {
					retChan <- workers[extension](f)
				}
			}
			waitGroup.Done()
		}()
	}
	go func() {
		waitGroup.Wait()
		close(retChan)
	}()

	// output
	for r := range retChan {
		fmt.Printf("file:%s total:%d empty:%d effective:%d comment:%d\n",
			r.filename, r.total, r.empty, r.effective, r.comment)
	}
}
