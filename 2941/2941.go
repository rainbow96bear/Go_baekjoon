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
	fmt.Fscanf(reader,"%s\n", &input)
	check := input+"  "
	answer := 0
	for i:=0 ; i<len(input) ; i++{
		switch string(check[i]){
		case "c" :
			if string(check[i+1])=="=" || string(check[i+1])=="-" {
				i++
			}
		case "d" :
			if string(check[i+1])=="z" && string(check[i+2])=="=" {
				i += 2
			}
			if string(check[i+1])=="-" {
				i++
			}
		case "l", "n" :
			if string(check[i+1])=="j" {
				i++
			}
		case "s", "z" :
			if string(check[i+1])=="=" {
				i++
			}
		}
		answer++
	}
	fmt.Fprint(writer,answer)
}