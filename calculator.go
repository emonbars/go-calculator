package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type calc struct{}

func (calc) operate(operator1 int, operator string, operator2 int) int {
	switch operator {
	case "+":
		return operator1 + operator2
	case "-":
		return operator1 - operator2
	case "*":
		return operator1 * operator2
	case "/":
		return operator1 / operator2
	default:
		return 0
	}
}

func readInput() string {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	return scaner.Text()
}

func parse(values string) int {
	operator, _ := strconv.Atoi(values)
	return operator
}

func main() {
	regex := regexp.MustCompile(`^(\d+)([\+\-\*\/])(\d+)$`)
	input := readInput()
	fmt.Println()
	if regex.Match([]byte(input)) {
		c := calc{}
		valueone := parse(regex.ReplaceAllString(input, "$1"))
		operation := regex.ReplaceAllString(input, "$2")
		valuetwo := parse(regex.ReplaceAllString(input, "$3"))
		fmt.Println(c.operate(valueone, operation, valuetwo))
	} else {
		fmt.Println("The input is not valid")
	}
}
