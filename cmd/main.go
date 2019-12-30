package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/locug/iniopt"
)

var (
	originalPath = flag.String("o", "", "location of the original file")
	currentPath  = flag.String("c", "", "location of the current file")
	outLocation  = flag.String("out", "", "location to put final sql file ")
)

func main() {
	flag.Parse()

	log.Println(*originalPath, *currentPath)

	if *originalPath == "" || *currentPath == "" {
		log.Panicln("both original and current path needed ")
	}

	b, err := iniopt.CompareINI(*originalPath, *currentPath)
	if err != nil {
		log.Panic(err)
	}

	name := filepath.Base(strings.TrimSuffix(*originalPath, filepath.Ext(*originalPath)))

	outPath := filepath.Join(*outLocation, name+".sql")

	err = ioutil.WriteFile(outPath, b, 0666)
	if err != nil {
		log.Panic(err)
	}
}
