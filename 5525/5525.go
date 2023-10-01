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

	var N int
	fmt.Fscan(reader, &N)

	var M int
	fmt.Fscan(reader, &M)

	var input string
	fmt.Fscan(reader, &input)

	answer := 0
	list := strings.Split(input, "")
	for i:=0 ; i<M ; i++ {
		if list[i] == "I" {
			cnt := 0
			for j:=1 ; i+(2*j) < M ; j++{
				if list[i+(2*j-1)] == "O" && list[i+(2*j)] == "I" {
					cnt++
				}else{
					break
				}
			}
			if cnt >= N {
				answer += (cnt-N+1)
			}
			i += cnt*2
		}
	}
	fmt.Fprintln(writer, answer)
}