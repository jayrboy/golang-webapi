package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func main() {
	e := employee{}

	// ใช้ json.Unmarshal และเพิ่ม e ตำแหน่ง Pointer ไปยังโครงสร้าง employee
	err := json.Unmarshal([]byte(`{"ID": 101, "Name": "Jay Jakkrit", "Tel":"0123456789","Email": "jakkrit@mail.com"}`), &e)
	if err != nil {
		log.Fatal("Error UnMarshaling JSON:", err)
		fmt.Println("Error UnMarshaling JSON:", err)
		return
	}
	fmt.Println(e)
	fmt.Println(e.Name)
}
