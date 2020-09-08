package main

import (
	"log"
	"os"
)

var (
	outFile *os.File
	LogFile *log.Logger
)

func WriteUpdateLog(update Update) {
	LogFile.Printf("%+v\n",update)
}

func WriteErrorLog(err error) {
	if err != nil {
		LogFile.Fatalln(err)
	}
}
