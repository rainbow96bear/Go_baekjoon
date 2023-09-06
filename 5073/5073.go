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

	for {
		m := make(map[int]int)
		side := make([]int,3)
		fmt.Fscan(reader, &side[0], &side[1], &side[2])
		sort.Ints(side)
		for i:=0 ; i<3 ;i++{
			m[side[i]]++
		}
		if len(m) == 1 {
			if side[0] == 0 {
				break
			}else{
				fmt.Fprintln(writer, "Equilateral")
			}
		}else if side[0]+side[1] <= side[2] {
			fmt.Fprintln(writer, "Invalid")
		}else {
			if len(m) == 2 {
				fmt.Fprintln(writer, "Isosceles")
			}else {
				fmt.Fprintln(writer, "Scalene")
			}
		}
	}
}