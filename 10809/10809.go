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

	var inputText string
	fmt.Fscan(reader, &inputText)
	var alphabet = make([]int,26)
	for i:=0 ; i<len(alphabet) ; i++{
		alphabet[i] = -1
	}
	for i:=0 ; i<len(inputText) ; i++{
		if alphabet[inputText[i]-'a'] == -1{
			alphabet[inputText[i]-'a'] = i
		}
	}
	for _, v := range alphabet{
		fmt.Fprintf(writer,"%d ",v)
	}
}