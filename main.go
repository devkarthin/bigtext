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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("----------------- Welcome to discord big text utility by karthin#0073 -----------------")

	for {

		fmt.Print("Enter text: ")
		if !scanner.Scan() {
			break
		}

		bigtext := bigify(scanner.Text())

		clipboard.WriteAll(bigtext)

		fmt.Println("----------------- Text copied! -----------------")

	}
}

func bigify(text string) string {
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
	return strings.Join(bigletters, " ")
}
