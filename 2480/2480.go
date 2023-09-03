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

	dice := make([]int,3)
	fmt.Fscan(reader, &dice[0], &dice[1], &dice[2])
	sort.Ints(dice)

	if dice[0] == dice[1] && dice[1] == dice[2]{
		fmt.Fprintln(writer, 10000+dice[2]*1000)
	}else if dice[0] == dice[1]{
		fmt.Fprintln(writer, 1000+dice[0]*100)
	}else if dice[1] == dice[2] {
		fmt.Fprintln(writer, 1000+dice[1]*100)
	}else {
		fmt.Fprintln(writer,dice[2]*100)
	}
}