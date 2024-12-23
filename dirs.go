package main

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func mkPaths(tree []knot) []string {
	paths := []string{}
	dirs := ""
	for _, k := range tree {
		dirs = k.name
		paths = append(paths, dirs) // todo filepath.Join
		if k.knot != nil {
			for _, v := range k.knot {
				paths = append(paths, filepath.Join(dirs, v.name))
			}
		}
	}
	return paths
}

func myMkdirs(path string) {
	if e := os.MkdirAll(path, 0755); e != nil {
		//log.Println(e)
	} // todo errors
}

// todo proper errors, or
// todo remove (for linx)
func myMkfile(name string) {
	if _, err := os.Stat(name); errors.Is(err, fs.ErrNotExist) {
		if f, e := os.Create(name); e != nil {
			log.Println("myMkfile can't create "+name, e)
			defer f.Close()
		}
	}
}
