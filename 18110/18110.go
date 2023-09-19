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

	var n int
	fmt.Fscan(reader, &n)
	if n == 0 {
		fmt.Fprintln(writer, 0)
		return
	}
	list := make([]int, n)
	for i:=0 ; i<n ; i++{
		fmt.Fscan(reader, &list[i])
	}
	sort.Ints(list)

	cut := halfUp(float64(n)*0.15)
	
	cuttedList := list[cut:(n-cut)]
	total := 0
	for i:=0 ; i<len(cuttedList) ; i++ {
		total += cuttedList[i]
	}
	
	answer := halfUp(float64(total)/float64(n-(2*cut)))
	fmt.Fprintln(writer, answer)
}



func halfUp(num float64)int{
	if int(num*10.0)%10 >= 5 {
		return int(num)+1
	}
	return int(num)
}