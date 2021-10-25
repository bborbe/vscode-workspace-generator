package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Folder struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path"`
}

type Config struct {
	Folders  []Folder `json:"folders"`
	Settings struct {
	} `json:"settings"`
}

func main() {
	var pathPtr = flag.String("path", "", "path")
	flag.Parse()
	if *pathPtr == "" {
		panic("param path missing")
	}

	var config Config

	for _, e := range createList(*pathPtr) {
		dir := strings.Replace(e, *pathPtr, "", 1)
		config.Folders = append(config.Folders, Folder{
			Name: dir,
			Path: dir,
		})
	}
	json.NewEncoder(os.Stdout).Encode(config)
}

func createList(dir string) []string {
	var result []string
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, e := range list {
		if e.IsDir() {
			if e.Name() == "vendor" {
				continue
			}
			result = append(result, createList(path.Join(dir, e.Name()))...)
		} else {
			if e.Name() == "go.mod" {
				result = append(result, dir)
			}
		}
	}
	return result
}
