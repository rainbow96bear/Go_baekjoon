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
	fmt.Fscanf(reader, "%d %d\n", &N, &M)
	order_list := make([][]int, M)
	for i:=0 ; i<M ; i++ {
		var length int
		fmt.Fscanf(reader, "%d", &length)
		order_list[i] = make([]int, length)
		for j:=0 ; j<length ; j++ {
			fmt.Fscanf(reader, "%d", &order_list[i][j])
		}
		fmt.Fscanf(reader, "\n")
	}

	result := Find_best_order(N, M, order_list)
	for _, num := range result {
		fmt.Fprintln(writer, num)
	}
}

func Find_best_order(N, M int, order_list [][]int) []int {
	value:= make(map[int]int)
	childList := make(map[int]map[int]bool)
	for i:=1 ; i<=N ; i++ {
		childList[i] = make(map[int]bool)
	}
	for i:=0 ; i<M ; i++ {
		for j:=0 ; j<len(order_list[i])-1 ; j++ {
			head := order_list[i][j]
			for k:=j+1 ; k<len(order_list[i]) ; k++ {
				if childList[head][order_list[i][k]] == false {
					childList[head][order_list[i][k]] = true
					value[order_list[i][k]]++
				}
			}
		}
	}

	result := []int{}
	for i:=1; i<=N ; i++ {
		if value[i] == 0 {
			result = append(result , i)
		}
	}
	for i:=0 ; i<len(result) ; i++ {
		num := result[i]
		for j:=1 ; j<=N ; j++ {
			if childList[num][j] == true {
				childList[num][j] = false
				value[j]--
				if value[j] == 0 {
					result = append(result, j)
				}
			}
		}
	}
	if len(result) == N {
		return result
	}
	return []int{0}
}