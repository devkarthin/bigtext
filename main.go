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

	for {
		fmt.Print("Enter text: ")
		if !scanner.Scan() {
			break
		}

		bigtext, err := bigify(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		if err := clipboard.WriteAll(bigtext); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Text copied")

	}
}

func bigify(text string) (string, error) {
	r, err := regexp.Compile("[\\s]+")
	if err != nil {
		return "", err
	}

	text = r.ReplaceAllString(text, " ")
	letters := strings.Split(text, "")
	var bigletters []string

	for _, l := range letters {
		emoji := getEmoji(l)
		bigletters = append(bigletters, emoji)

	}
	return strings.Join(bigletters, " "), nil
}
