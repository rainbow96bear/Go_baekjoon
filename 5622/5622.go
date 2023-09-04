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

	s,_ := reader.ReadString('\n')
	var answer int = 0
	for _,v := range s{
		switch string(v) {
		case "A", "B", "C" :
			answer += 3
		case "D", "E", "F" :
			answer += 4
		case "G", "H", "I" :
			answer += 5
		case "J", "K", "L" :
			answer += 6
		case "M", "N", "O" :
			answer += 7
		case "P", "Q", "R", "S" :
			answer += 8
		case "T", "U", "V" :
			answer += 9
		case "W", "X", "Y", "Z" :
			answer += 10
		}
	}
	fmt.Fprint(writer, answer)
}