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

	sli := make([]int,5)
	sum := 0
	for i:=0 ; i<5 ; i++{
		fmt.Fscan(reader, &sli[i])
		sum += sli[i]
	}
	sort.Ints(sli)
	fmt.Fprintf(writer, "%d\n%d",sum/5,sli[2])

}