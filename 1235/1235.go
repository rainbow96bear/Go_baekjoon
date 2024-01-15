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

	list := make([]string, N)
	for n:=0 ; n<N ; n++ {
		fmt.Fscanf(reader, "%s\n", &list[n])
	}
	result := FindMinLength(N, list)
	fmt.Fprintln(writer, result)
}

func FindMinLength(N int, list []string) int {
	for i:=len(list[0])-1 ; i>=0 ; i-- {
		check := make(map[string]bool)
		isPossible := true
		for j:=0 ; j<len(list) ; j++ {
			if check[list[j][i:]] == false {
				check[list[j][i:]] = true
			}else {
				isPossible = false
				break
			}
		}
		if isPossible {
			return len(list[0]) - i
		}
	}
	return len(list[0])
}