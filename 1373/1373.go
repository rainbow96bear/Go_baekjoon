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

	var (
		input string
	)
	fmt.Fscanf(reader, "%s\n", &input)
	
	if len(input)%3 == 1 {
		input = "00"+input
	}else if len(input)%3 == 2 {
		input = "0" + input
	}

	for i:=0 ; 3*i<len(input) ; i++ {
		fmt.Fprintf(writer, "%d",(input[3*i]-'0')*4 + (input[3*i+1]-'0')*2 + (input[3*i+2]-'0')) 
	}
}
