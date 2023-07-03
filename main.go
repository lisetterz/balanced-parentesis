package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the balanced parenthesis validator.\n Please enter the string to validate:")

	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')

	word = strings.Replace(word, "\n", "", -1)

	if err != nil {
		panic(err)
	}

	match := validateParentheses(word)

	if match {
		fmt.Println("The string is valid")
	} else {
		fmt.Println("The string is invalid")
	}

}

func validateParentheses(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {

		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		st = st[:l]
	}

	return len(st) == 0
}
