package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()
	
	arr := []int{}
	for scanner.Scan() {
        input := scanner.Text()
        if input == "" {
            break
        }
        value, _ := strconv.Atoi(input)
        arr = append(arr, value)
    }
	result := Process(arr)
	for _,v := range result {
		fmt.Println(v)
	}
}
func Process(arr []int) []int{
	if len(arr) == 0 {
		return []int{}
	}
	Parent := arr[0]
	left := []int{}
	right := []int{}
	for i:=1 ; i<len(arr) ; i++ {
		if Parent > arr[i] {
			left = append(left, arr[i])
		}else {
			right = append(right, arr[i])
		}
	}
	left = Process(left)
	right = Process(right)
	result := left
	result = append(result, right...)
	result = append(result, Parent)
	return result
}