package main

import (
	"fmt"
	"math"
	"strconv"
)

var n, m int
var ans = -1

func isSquareNumber(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n, &m)
	arr := make([][]int, n+1)
	for i := range arr {
		arr[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		for j := 1; j <= m; j++ {
			arr[i][j], _ = strconv.Atoi(string(tmp[j-1]))
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for x := -n; x < n; x++ {
				for y := -m; y < m; y++ {
					if x == 0 && y == 0 {
						continue
					}
					a, b := i, j
					now := 0
					for a > 0 && a <= n && b > 0 && b <= m {
						now *= 10
						now += arr[a][b]
						if isSquareNumber(now) {
							ans = max(ans, now)
						}
						a += x
						b += y
					}
					if isSquareNumber(now) {
						ans = max(ans, now)
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
