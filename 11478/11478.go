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

	var text string
	fmt.Fscan(reader, &text)

	list := make(map[string]bool)
	for i:=0 ; i<len(text) ; i++{
		for j:=i+1 ; j<=len(text) ; j++{
			word := text[i:j]
			list[word] = true
		}
	}
	fmt.Fprint(writer,len(list))
}