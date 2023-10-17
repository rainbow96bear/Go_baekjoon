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

	arr := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	
	result := FindLength(arr)
	fmt.Fprintln(writer, result)
}
func FindLength(arr []int) int {
	check := make(map[int]int)
	max := 0 
	for i:=1 ; i<len(arr) ; i++ {
		for j:=0 ; j<i ; j++{
			if arr[j] < arr[i] {
				if check[i] < check[j]+1 {
					check[i] = check[j]+1
				}
			}
		}
		if check[i] > max {
			max = check[i]
		}
	}
	return max + 1
}