package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	arr := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {return arr[i]<arr[j]})
	answer := Process(arr, M)
	for i:=0 ; i<len(answer) ; i++ {
		fmt.Fprintln(writer, strings.Join(answer[i], " "))
	}
}

func Process(arr []int, cnt int) [][]string {
	if cnt == 0 {
		return [][]string{}
	}
	isUse := make(map[int]bool)
	result := [][]string{}
	for i:=0 ; i<len(arr) ; i++ {
		if isUse[arr[i]] == false {
			list := Process(arr[i:], cnt-1)
			for j:=0 ; j<len(list) ; j++ {
				newElement := []string{fmt.Sprint(arr[i])}
				newElement = append(newElement, list[j]...)
				result = append(result, newElement)
			}
			if len(list) == 0 {
				result = append(result,[]string{fmt.Sprint(arr[i])})
			}
			isUse[arr[i]]=true
		}
	}
	return result
}