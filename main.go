package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------- Welcome to discord big text utility by karthin#0073 -----------------")
	fmt.Print("Enter text: ")

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	r, _ := regexp.Compile("\\s+")
	text = r.ReplaceAllString(text, " ")

	letters := strings.Split(text, "")

	var bigletters []string

	for _, value := range letters {
		if value != " " {
			bigletters = append(bigletters, ":regional_indicator_"+strings.ToLower(value)+":")
		} else {
			bigletters = append(bigletters, " ")
		}
	}
	// fmt.Println("----------------- Here ya go -----------------")
	// fmt.Println(strings.Join(bigletters, " "))
	// fmt.Println("----------------------------------------------")
	bigtext := strings.Join(bigletters, " ")

	clipboard.WriteAll(bigtext)

	fmt.Println("----------------- Text copied! -----------------")

}
