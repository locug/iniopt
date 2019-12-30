package iniopt

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReadFiles(t *testing.T) {
	b, err := CompareINI("./original.ini", "current.ini")
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}

	fmt.Print(string(b))

	ioutil.WriteFile("testout.sql", b, 0666)

}
