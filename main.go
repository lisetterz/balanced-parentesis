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

func validateParentheses(str string) bool {

	if len(str)%2 != 0 {
		return false
	}

	st := []rune{}
	parentheses := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, j := range str {

		if closer, ok := parentheses[j]; ok {
			st = append(st, closer)
			continue
		}

		length := len(st) - 1
		if length < 0 || j != st[length] {
			return false
		}

		st = st[:length]
	}

	return len(st) == 0
}
