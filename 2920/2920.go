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

	soundList := make([]int,8)
	for i:=0 ; i<8 ; i++ {
		fmt.Fscan(reader, &soundList[i])
	}

	for i:=0 ; i<8 ; i++ {
		if soundList[i] == i+1 && answer != "descending" {
			answer = "ascending"
		}else if soundList[i] == 8-i && answer != "ascending" {
			answer = "descending"
		} else {
			answer = "mixed"
			break
		}
	}
	fmt.Fprint(writer, answer)
}