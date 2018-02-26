package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func activate() {
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
		fmt.Printf("Activated %s...\n", newPath)
		err := os.Symlink(f, newPath)

		// Continue running if file already existed
		fileExistsErrorMsg := fmt.Sprintf("symlink %s %s: file exists", f, newPath)
		if err != nil && err.Error() != fileExistsErrorMsg {
			fmt.Println(err.Error())
			log.Fatal(err)
		}
	}
}

func deactivate() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	swampPath := path.Join(u.HomeDir, ".swamp")

	// Remove all the symlinks
	files := []string{}
	filepath.Walk(swampPath, func(p string, f os.FileInfo, err error) error {
		files = append(files, p)
		return nil
	})

	for _, f := range files[1:] {
		symPath := path.Join(u.HomeDir, path.Base(f))
		fmt.Printf("Deactivating %s...\n", symPath)
		err = os.Remove(symPath)
		if err != nil {
			fmt.Printf("File did not exist: %s\n", symPath)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "swamp"
	app.Usage = "organize and share your configuration files with ease"
	app.Commands = []cli.Command{
		{
			Name:    "activate",
			Aliases: []string{"a"},
			Usage:   "activate your swamp",
			Action: func(c *cli.Context) error {
				activate()
				return nil
			},
		},
		{
			Name:    "deactivate",
			Aliases: []string{"a"},
			Usage:   "deactivate your swamp",
			Action: func(c *cli.Context) error {
				deactivate()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
