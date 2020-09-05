package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

const START_LEVEL = 0

type treePrintDTO struct {
	output       *io.Writer
	path         string
	includeFiles bool
}

type filePrintDTO struct {
	output *io.Writer
	file   *os.FileInfo
	isLast bool
	level  int
	isLastDir bool
}

func checkIsPosLast(files []os.FileInfo, pos int) bool  {
	for i:=pos+1; i < len(files) ; i++ {
		if files[i].IsDir() {
			return false
		}
	}
	return true
}

func printFile(dto *filePrintDTO) {
	file := *dto.file
	var prefix string

	if dto.isLast {
		prefix = "└───"
	} else {
		prefix = "├───"
	}

	if dto.level == START_LEVEL {
		fmt.Fprintf(*dto.output, prefix + file.Name() + "\n")
		return
	}

	indent := "|\t"
	lastDirFlag := dto.isLastDir
	for i:=0; i < dto.level; i++ {
		if lastDirFlag {
			fmt.Fprintf(*dto.output, indent)
			indent = "\t"
			lastDirFlag = false
			continue
		}
		fmt.Fprintf(*dto.output, indent)
	}

	fmt.Fprintf(*dto.output, prefix + file.Name() + "\n")

}

func newTreePrintDTO(output *io.Writer, path string, includeFiles bool) *treePrintDTO {
	return &treePrintDTO{output, path, includeFiles}
}

func printTree(dto *treePrintDTO, level int) error {
	file, err := os.Open(dto.path)

	if err != nil {
		return err
	}

	files, err := file.Readdir(0)

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	tmpPath := dto.path

	isLastDir := false
	for pos, file := range files {
		if file.IsDir() {
			isLast := checkIsPosLast(files, pos) && !dto.includeFiles
			isLastDir = isLast
			dto.path += "/" + file.Name()
			printDto := &filePrintDTO{dto.output, &file, isLast, level,isLastDir}
			printFile(printDto)
			level++
			if err := printTree(dto, level); err != nil {
				return err
			}
			isLastDir = false
			level--
			dto.path = tmpPath

		} else {
			if dto.includeFiles {
				isLast := pos == len(files) - 1
				printDto := &filePrintDTO{dto.output, &file, isLast, level,isLastDir}
				printFile(printDto)
			}
		}

	}

	return nil
}

func dirTree(output io.Writer, path string, includeFiles bool) error {
	dto := newTreePrintDTO(&output, path, includeFiles)

	return printTree(dto,START_LEVEL)
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
