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

	var L, R string
	fmt.Fscan(reader, &L, &R)

	leftSlice := strings.Split(L, "")
	rightSlice := strings.Split(R, "")
	if len(leftSlice) != len(rightSlice){
		fmt.Fprint(writer, 0)
	}else {
		cnt := 0
		for i:=0 ; i<len(leftSlice) ; i++{
			if leftSlice[i] != rightSlice[i] {
				break
			}else if leftSlice[i] == "8" && rightSlice[i] == "8" {
				cnt++
			}
		}
		fmt.Fprint(writer, cnt)
	}
}