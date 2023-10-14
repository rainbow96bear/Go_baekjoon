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
	fmt.Fscan(reader, &N, &M)

	arr := make([]int, M+1)
	for i:=0 ; i<=M ; i++ {
		arr[i]=1
	}
	for arr[0] < 2 {
		for i:=1 ; i<=M ; i++ {
			fmt.Fprintf(writer, "%d ", arr[i])
		}
		fmt.Fprintln(writer)
		arr[M]++
		for i:=0 ; i<M ; i++ {
			if arr[M-i] > N {
				arr[M-i-1]++
				for j:=0 ; j<=i ; j++ {
					arr[M-i+j] = arr[M-i+j-1]
				}
			}
		}
	}
}