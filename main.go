package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	version  = "v0.1"
	excludes = []string{
		".jekyll-metadata",
		".sass-cache",
		"_site",
		"tmp",
	}
	logger = log.New(ioutil.Discard, "", log.LstdFlags)
)

func hasBomb(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		logger.Printf("failed to open: %s", err)
		return false
	}
	defer f.Close()
	br := bufio.NewReader(f)
	r, _, err := br.ReadRune()
	if err != nil {
		logger.Printf("failed to read: %s", err)
		return false
	}
	return r == '\uFEFF'
}

func main() {
	verbose := flag.Bool("verbose", false, "verbose messages")
	version := flag.Bool("version", false, "show version number")
	flag.Parse()
	if *version {
		fmt.Printf("detectbomb version %s\n", version)
		os.Exit(1)
	}
	if *verbose {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	ignores := make(map[string]bool)
	for _, n := range excludes {
		ignores[n] = true
	}
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		_, ok := ignores[info.Name()]
		if ok {
			logger.Printf("SKIPPED: %s", path)
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if hasBomb(path) {
			fmt.Printf("BOMB: %s\n", path)
		}
		return nil
	})
}
