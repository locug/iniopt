package main

import "github.com/locug/iniopt"

func main() {
	iniopt.ReadFiles("./original.ini", "./current.ini")
}
