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

		bigtext := bigify(strings.ToLower(scanner.Text()))

		clipboard.WriteAll(bigtext)

		fmt.Println("----------------- Text copied! -----------------")

	}
}

func bigify(text string) string {

	// r, _ := regexp.Compile("\\s+")
	text = regexp.MustCompile("[\\s]+").ReplaceAllString(text, " ")

	letters := strings.Split(text, "")

	var bigletters []string

	for index, value := range letters {
		if regexp.MustCompile("[a-zA-Z]").MatchString(value) {
			value = ":regional_indicator_" + value + ":"
		} else if regexp.MustCompile("[0-9]").MatchString(value) {
			switch value {
			case "0":
				value = ":zero:"
			case "1":
				value = ":one:"
			case "2":
				value = ":two:"
			case "3":
				value = ":three:"
			case "4":
				value = ":four:"
			case "5":
				value = ":five:"
			case "6":
				value = ":six:"
			case "7":
				value = ":seven:"
			case "8":
				value = ":eight:"
			case "9":
				value = ":nine:"
			case "10":
				value = ":keycap_ten:"
			}
		} else if value == " " && letters[index+1] != "$" {
		} else if value == "#" {
			value = ":hash:"
		} else if value == "?" {
			value = ":question:"
		} else {
			bigletters = append(bigletters, value)
			continue
		}
		bigletters = append(bigletters, value+" ")
	}
	// fmt.Println("----------------- Here ya go -----------------")
	// fmt.Println(strings.Join(bigletters, " "))
	// fmt.Println("----------------------------------------------")
	return strings.Join(bigletters, "")
}
