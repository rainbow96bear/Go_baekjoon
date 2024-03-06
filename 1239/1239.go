package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	chart    []int    
	check    []bool   
	result   int      
	N        int      
	fullMask int      
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scanln(&N)

	var max int = -1
	chart = make([]int, N)
	check = make([]bool, 51)
	fullMask = (1 << N) - 1

	inputLine := ""
	if scanner.Scan() {
		inputLine = scanner.Text()
	}

	values := strings.Fields(inputLine)
	for i := 0; i < N; i++ {
		chart[i], _ = strconv.Atoi(values[i])
		max = findMax(max, chart[i])
	}

	if max > 50 {
		fmt.Println("0")
	} else if max == 0 {
		fmt.Println("1")
	} else {
		check[0] = true
		search(0, 0, false, 0)
		fmt.Println(result)
	}
}

func search(cnt, sum int, flag bool, mask int) {
	if mask == fullMask {
		result = findMax(result, cnt)
		return
	}

	for i := 0; i < N; i++ {
		if (mask & (1 << i)) >= 1 {
			continue
		}
		temp := chart[i] + sum

		if !flag {
			if temp >= 50 {
				if check[temp%50] {
					search(cnt+1, temp%50, true, mask|(1<<i))
				} else {
					search(cnt, temp%50, true, mask|(1<<i))
				}
			} else {
				check[temp] = true
				search(cnt, temp, false, mask|(1<<i))
				check[temp] = false
			}
		} else {
			if check[temp] {
				search(cnt+1, temp, true, mask|(1<<i))
			} else {
				search(cnt, temp, true, mask|(1<<i))
			}
		}
	}
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
