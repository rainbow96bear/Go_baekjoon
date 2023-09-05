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

	var N, M int
	fmt.Fscan(reader, &M, &N)
	min := -1
	sum := 0
	for i:=M ; i<=N ; i++{
		check := false
		if i == 1 {
			continue
		}
		if i == 2 {
			min = i
			sum += i
			continue
		}
		for j:=2 ; j<i ; j++{
			if j != i && i%j == 0{
				check = true
				break
			}
		}
		if check {
			continue
		}
		if sum == 0 {
			min = i
		}
		sum += i
	}
	if min == -1 {
		fmt.Fprint(writer, min)
	}else {
		fmt.Fprint(writer, sum, min)
	}
}