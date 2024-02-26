package main

import (
	"fmt"
	"sort"
)

var n, k, tmp int
var numarr, resultarr []int

func compare(lhs, rhs int) bool {
	return lhs > rhs
}

func makecase(num, cnt int) {
	if cnt != 0 && num <= n {
		resultarr = append(resultarr, num)
	}
	if num > n {
		return
	}
	for i := 0; i < k; i++ {
		if cnt == 0 {
			makecase(num*numarr[i], cnt+1)
		} else {
			makecase(num*10+numarr[i], cnt+1)
		}
	}
}

func main() {
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		fmt.Scan(&tmp)
		numarr = append(numarr, tmp)
	}
	sort.Slice(numarr, func(i, j int) bool {
		return compare(numarr[i], numarr[j])
	})

	makecase(1, 0)
	sort.Slice(resultarr, func(i, j int) bool {
		return compare(resultarr[i], resultarr[j])
	})
	fmt.Println(resultarr[0])
}
