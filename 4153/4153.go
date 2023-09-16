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
		sides := make([]int,3)
		fmt.Fscan(reader, &sides[0], &sides[1], &sides[2])
		if sides[0] == 0 && sides[1] == 0 && sides[2] == 0 {
			return
		}
		sort.Ints(sides)
		sqrt1 := sides[0]*sides[0]
		sqrt2 := sides[1]*sides[1]
		sqrt3 := sides[2]*sides[2]
		if sqrt3 == sqrt1 + sqrt2 {
			fmt.Fprint(writer, "right\n")
		}else {
			fmt.Fprint(writer, "wrong\n")
		}
	}
}