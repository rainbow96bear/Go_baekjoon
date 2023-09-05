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
	var testCase int
	fmt.Fscan(reader, &testCase)
	answer := testCase
	for i:=0 ; i<testCase ; i++{
		var input string
		fmt.Fscan(reader, &input)

		checkChar := make(map[byte]bool)
		preChar := input[0]
		checkChar[preChar] = true
		for i:=1 ; i<len(input) ; i++{
			if preChar != input[i]{
				if checkChar[input[i]]{
					answer-- 
					break
				}
				preChar=input[i]
				checkChar[input[i]]=true
			}
		}
	}
	fmt.Fprint(writer, answer)
}