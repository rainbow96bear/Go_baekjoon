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

	var N int
	fmt.Fscan(reader, &N)
	
	computerList := make(map[int][]int)

	var M int
	fmt.Fscan(reader, &M)

	check := make(map[int]bool)
	for i:=0 ; i<M ; i++ {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		computerList[A] = append(computerList[A], B)
		computerList[B] = append(computerList[B], A)
	}

	stack := []int{1}
	check[1]=true
	cnt := 0
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i:=0 ; i<len(computerList[pop]) ; i++ {
			if check[computerList[pop][i]] == false {
				stack = append(stack, computerList[pop][i])
				check[computerList[pop][i]] = true
				cnt++
			}
		}
	}
	fmt.Fprintln(writer, cnt)
}