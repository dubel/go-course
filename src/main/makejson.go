package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Write a program which prompts the user to first enter a name, and then enter an address.
//
//	Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
//	Your program should use Marshal() to create a JSON object from the map,
//
// and then your program should print the JSON object.
func main() {
	var nameInput, addressInput string
	println("Input name: ")
	_, err := fmt.Scan(&nameInput)
	if err != nil {
		log.Fatal(err)
	}
	println("Input address: ")
	reader := bufio.NewReader(os.Stdin)
	addressInput, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var nameMap = map[string]string{
		"name":    nameInput,
		"address": addressInput,
	}
	a, _ := json.Marshal(nameMap)
	fmt.Println(string(a))
}
