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

	fmt.Fscanf(reader, "%d\n", &N)
	
	result := Process(N)

	fmt.Fprintln(writer, result)
}

func Process(N int) int {
	check := make(map[int]bool)
	list := []int{}
	for i:=2 ; i<=N ; i++ {
		if check[i] == false {
			list = append(list, i)
			for num := i ; num<=N ; num+=i {
				check[num] = true
			}
		}
	}
	start, end := 0, 0
	sum := 0
	count := 0
	for start < len(list) {
		if sum > N {
			sum -= list[start]
			start++
		}else {
			if sum == N {
				count++
			}
			sum += list[end]
			if end != len(list)-1 {	
				end++
			}
		}
	}
	return count
}