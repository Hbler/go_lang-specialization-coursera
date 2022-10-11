package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	ReadFile()
}

type Name struct {
	fname string
	lname string
}

func getFileName() string {
	var filename string

	name := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a name:")
	name.Scan()
	filename = name.Text()

	return string(filename)
}

func ReadFile() {
	filename := getFileName()
	absPath, _ := filepath.Abs("./read")
	file, err := os.Open(absPath + "/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	namesSlice := make([]Name, 0, 3)

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		aName := Name{
			strings.Split(reader.Text(), " ")[0],
			strings.Split(reader.Text(), " ")[1],
		}
		namesSlice = append(namesSlice, aName)

	}

	for _, v := range namesSlice {
		fmt.Printf("%s %s\n", v.fname, v.lname)
	}
}
