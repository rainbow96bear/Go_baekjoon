package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()

	for {
		var input string
		input, _ = reader.ReadString('\n')
		if input[0] == '.'{
			return
		}
		input = strings.TrimSpace(input)
		stack := []rune{}

		for _, char := range input {
			switch char {
			case '(', '[':
				stack = append(stack, char)
			case ')':
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					stack = append(stack, char)
				} else {
					stack = stack[:len(stack)-1]
				}
			case ']':
				if len(stack) == 0 || stack[len(stack)-1] != '[' {
					stack = append(stack, char)
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if len(stack) > 0 {
			fmt.Fprintln(writer, "no")
		}else {
			fmt.Fprintln(writer, "yes")
		}
	}
}