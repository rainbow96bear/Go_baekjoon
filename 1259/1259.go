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

	for {
		var input string
		fmt.Fscan(reader, &input)
		if input == "0" {
			return
		}
		check := true
		for i, v := range input {
			if string(v) != string(input[len(input)-1-i]) {
				fmt.Fprintln(writer, "no")
				check = false
				break
			}
		}
		if check {
			fmt.Fprintln(writer, "yes")
		}
	}
}