package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	
	list := make([]int,N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &list[i])
	}
	sort.Ints(list)

	var M int
	fmt.Fscan(reader, &M)

	for i:=0 ; i<M ; i++ {
		var input int
		fmt.Fscan(reader, &input)

		min := 0
		max := N-1
		check := true
		for min <= max {
			mid := (min+max) / 2
			if list[mid] < input {
				min = mid+1
			}else if list[mid] > input {
				max = mid-1
			}else {
				fmt.Fprintln(writer, 1)
				check = false
				break
			}
		}
		if check {
			fmt.Fprintln(writer, 0)
		}
	}
}
