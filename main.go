package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Dumb Lisp!")
	for {
		fmt.Print("\nλ ")
		r := bufio.NewReader(os.Stdin)
		input, _ := r.ReadString('\n')

		if input == "(exit)\n" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		var stack nestStack

		for i, r := range input {
			switch r {
			case '(':
				stack.Push(i)

			case ')':
				iOpen, err := stack.Pop()
				if err != nil {
					log.Fatal(err)
				}
				innerSexp := input[iOpen : i+1]
				eval(innerSexp)
			}
		}
	}
}

func eval(sexp string) {
	fmt.Println(sexp)
}

type nestStack []int

func (s *nestStack) Push(val int) {
	*s = append(*s, val)
}

func (s *nestStack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("Empty stack")
	}

	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return popped, nil
}
