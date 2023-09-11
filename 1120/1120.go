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

	var A, B string
	fmt.Fscan(reader, &A, &B)
	Asplit := strings.Split(A,"")
	Bsplit := strings.Split(B,"")
	answer := len(Bsplit)
	for i:=0 ; i<= len(Bsplit)-len(Asplit) ; i++ {
		cnt := 0
		for j:=i ; j-i<len(Asplit) ; j++ {
			if Asplit[j-i] !=  Bsplit[j] {
				cnt++
			}
		}
		if cnt < answer {
			answer = cnt
		}
	}
	fmt.Fprint(writer, answer)
}