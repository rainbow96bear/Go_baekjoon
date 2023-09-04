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
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Fprint(writer,input)
	}
}