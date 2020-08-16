package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func printTree(output *io.Writer, path *string, includeFiles *bool) error {
	file, err := os.Open(*path)
	if err != nil {
		return err
	}

	files, err := file.Readdir(0)

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	if len(files) == 1 {
		fmt.Fprintf(*output,"|\t")
	}

	for pos,file := range files {
		if file.IsDir() {
			fmt.Fprintf(*output, "├───"+file.Name()+"\n")
			tmpPath := *path + "/"+file.Name()
			if err := printTree(output, &tmpPath, includeFiles); err != nil {
				return err
			}

		} else {
			if *includeFiles {
				if pos == len(files)-1{
					fmt.Fprintf(*output,"|\t")
					fmt.Fprintf(*output,"└───")
				} else {
					fmt.Fprintf(*output,"├───")
				}
				fmt.Fprintf(*output,file.Name() +"\n")
			}
		}

	}
	return nil
}

func dirTree(output io.Writer, path string, includeFiles bool ) error {
	printTree(&output, &path, &includeFiles)

   	return nil
}

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
