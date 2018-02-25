package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	swampPath := path.Join(u.HomeDir, ".swamp")

	// Make the swamp directory
	if _, err := os.Stat(swampPath); os.IsNotExist(err) {
		err = os.MkdirAll(swampPath, 0755)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Println("Created the directory", swampPath)
	}

	// Create symbolic links of all the files in this directory
	files := []string{}
	filepath.Walk(swampPath, func(p string, f os.FileInfo, err error) error {
		files = append(files, p)
		return nil
	})

	for _, f := range files[1:] {
		newPath := path.Join(u.HomeDir, path.Base(f))
		err := os.Symlink(f, newPath)
		if err != nil {
			log.Fatal()
		}
	}
}
