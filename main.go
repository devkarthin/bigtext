package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

var welcomemsg = strings.Repeat("-", 5) + " Welcome to big text utility v2 by karthin#0073 " + strings.Repeat("-", 5)
var copiedmsg = strings.Repeat("-", 10) + " Text copied " + strings.Repeat("-", 10)

func main() {

	fmt.Println(welcomemsg)
	fmt.Println("DISCLAIMER: THIS IS A PRE-PRODUCTION RELEASE. PLEASE REPORT ANY BUGS/ISSUES TO MY GITHUB OR BY CONTACTING ME ON DISCORD")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	// infinite loop of asking for input
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

		fmt.Println(copiedmsg)

	}
}

func bigify(text string) (string, error) {
	// regex check for any excess whitespace
	r, err := regexp.Compile("[\\s]+")
	if err != nil {
		return "", err
	}

	// removes any excess whitespace
	text = r.ReplaceAllString(text, " ")

	// splits the text letter by letter into an array
	letters := strings.Split(text, "")
	var bigletters []string

	// trys to get an emoji for each letter
	// if an emoji cannot be found, then skip it
	for _, l := range letters {
		emoji, _ := getEmoji(l)
		bigletters = append(bigletters, emoji)

	}
	return strings.Join(bigletters, " "), nil
}
