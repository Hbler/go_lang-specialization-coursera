package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	MakeJSON()
}

type info string

var m map[string]info

func MakeJSON() {
	m = make(map[string]info)
	name := bufio.NewScanner(os.Stdin)
	address := bufio.NewScanner(os.Stdin)

	fmt.Println("Type a name:")
	name.Scan()
	m["name"] = info(name.Text())

	fmt.Println("Type an address:")
	address.Scan()
	m["address"] = info(address.Text())

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("This is the map marshalled to JSON:")
	os.Stdout.Write(b)

}
