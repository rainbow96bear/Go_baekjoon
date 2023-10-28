package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	answer := ""
	for i:=0 ; i<15 ; i++ {
		input := Input(scanner)
		for _, v := range input {
			if v == 'w' {
				answer = "chunbae"
			}else if v == 'b'{
				answer = "nabi"
			}else if v == 'g' {
				answer = "yeongcheol"
			}
		}
	}
	fmt.Fprintln(writer, answer)
}


func Input(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()

	return input
}