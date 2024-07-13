package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Jay Jakkrit", "0123456789", "jakkrit@mail.com"})
	fmt.Println(string(data))
}
