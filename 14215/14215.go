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

	side := make([]int,3)
	fmt.Fscan(reader, &side[0], &side[1], &side[2])
	sort.Ints(side)
	if side[0]+side[1] <= side[2]{
		fmt.Fprint(writer, (side[0]+side[1])*2-1)
	}else{
		fmt.Fprint(writer, side[0]+side[1]+side[2])
	}
}