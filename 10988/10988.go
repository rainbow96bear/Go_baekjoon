package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	s,_ := reader.ReadString('\n')
	s=strings.TrimSpace(s)
	strSlice := strings.Split(s,"")

	answer := 1

	for i,v := range strSlice{
		if v != string(strSlice[len(strSlice)-1-i]){
			answer = 0
		}
	}
	fmt.Fprint(writer, answer)
}