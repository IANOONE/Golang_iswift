package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcword() {
// input
fmt.Print("Enter keyword : ")
reader := bufio.NewReader(os.Stdin)

input,_ := reader.ReadString('\n')

input = strings.TrimSuffix(input, "\r\n")

// split string into slice
entries  := strings.Split(input, "")

// declare map
result := map[string]int{}

for _, entry := range entries {
	value, exist := result[entry]
	if(exist){
		result[entry] = value + 1
	} else {
		result[entry] = 1
	}

	fmt.Println(entry)
}
fmt.Println(result)
}

func main() {
	calcword()
}