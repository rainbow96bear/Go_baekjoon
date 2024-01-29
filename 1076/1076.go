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
	answer := ""
	for i:=0 ; i<3 ; i++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)
		if i != 2 {
			if ChangeToValue(input) != "0" || i != 0 {
				answer += ChangeToValue(input)
			}
		}else {
			if answer != "0" {
				answer += ChangeToMultiple(input)
			}
		}
	}
	fmt.Fprintln(writer, answer)
}

func ChangeToValue(color string) (value string){
	switch color {
		case "black" :
			value = "0"
		case "brown" :
			value = "1"
		case "red" :
			value = "2"
		case "orange" :
			value = "3"
		case "yellow" :
			value = "4"
		case "green" :
			value = "5"
		case "blue" :
			value = "6"
		case "violet" :
			value = "7"
		case "grey" :
			value = "8"
		case "white" :
			value = "9"
	}
	return value
}

func ChangeToMultiple(color string) (value string) {
	switch color {
		case "black" :
			value = ""
		case "brown" :
			value = "0"
		case "red" :
			value = "00"
		case "orange" :
			value = "000"
		case "yellow" :
			value = "0000"
		case "green" :
			value = "00000"
		case "blue" :
			value = "000000"
		case "violet" :
			value = "0000000"
		case "grey" :
			value = "00000000"
		case "white" :
			value = "000000000"
	}
	return value
}