package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	for {
		var input string
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = input[:len(input)-1] // 개행 문자 하나만 제거

		

		count := 0
		for _, v := range input {
			if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
				count++
			}
			if v == '#' {
				return
			}
		}
		fmt.Fprintln(writer, count)
	}
}
