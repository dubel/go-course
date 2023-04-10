package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please input String to be parsed:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var strInLower = strings.ToLower(line)
	var trimmedString = strings.TrimSpace(strInLower)
	if strings.HasPrefix(trimmedString, "i") && strings.HasSuffix(trimmedString, "n") && strings.Contains(trimmedString, "a") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
