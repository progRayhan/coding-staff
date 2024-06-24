package main

import (
	"fmt"
	"go_staff/functions"
	"go_staff/structs"
)

func main() {
	myAdd := structs.Address{}
	myAdd2 := functions.MyAddress(myAdd)
	fmt.Println(myAdd2.Name)
}
