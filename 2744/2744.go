package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var input string
	fmt.Fscanf(reader, "%s\n", &input)

	result := changeBigSmall(input)
	fmt.Fprintln(writer, result)
}

func changeBigSmall(input string) string {
	result := ""
	for _, str := range input {
		if 'a' <= str && str <= 'z' {
			str += 'A'-'a'
			result += string(str)
		}else {
			str -= 'A'-'a'
			result += string(str)
		}
	}
	return result
}