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

	list := make([]int, 9)
	total := 0
	for i, _ := range list {
		fmt.Fscanf(reader, "%d\n", &list[i])
		total += list[i]
	}
	answer := make([]int, 7)
	isFind := false
	for i:=0 ; i<9 ; i++ {
		for j:=i+1 ; j<9 ; j++ {
			if total - list[i] - list[j] == 100 {
				answer = list[0:i]
				answer = append(answer, list[i+1:j]...)
				answer = append(answer, list[j+1:9]...)
				sort.Slice(answer, func(i, j int)bool{return answer[i] < answer[j]})
				
				isFind = true
				break
			}
		}
		if isFind {
			break
		}
	}

	for _, v := range answer {
		fmt.Fprintln(writer, v)
	}
}