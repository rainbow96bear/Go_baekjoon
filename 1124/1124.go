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

	var A, B int
	fmt.Fscan(reader, &A, &B)

	list := make([]int,B+1)
	for i:=2 ; 2*i<=B ; i++ {
		if list[i] == 0 {
			for j:=1 ; i*j<=B ; j++{
				k:= i*j
				for k%i == 0{
					list[i*j]++
					k/=i
				}
			}
		}
	}

	answer := 0
	for i:=A ; i<=B ; i++{
		if list[list[i]] == 1{
			answer++
		}
	}
	fmt.Fprint(writer, answer)
}