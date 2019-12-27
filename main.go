package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------- Welcome to discord big text utility by karthin#0073 -----------------")
	fmt.Print("Enter text: ")

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	letters := strings.Split(text, "")

	var bigtext []string

	for _, value := range letters {
		if value != " " {
			bigtext = append(bigtext, ":regional_indicator_"+strings.ToLower(value)+":")
		} else {
			bigtext = append(bigtext, " ")
		}
	}
	// fmt.Println("----------------- Here ya go -----------------")
	// fmt.Println(strings.Join(bigtext, " "))
	// fmt.Println("----------------------------------------------")
	clipboard.WriteAll(strings.Join(bigtext, " "))

	fmt.Println("----------------- Text copied! -----------------")

}
