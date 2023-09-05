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

	var x, y, w, h int
	fmt.Fscan(reader, &x, &y, &w, &h)
	check := []int{x, y, w-x, h-y}
	sort.Ints(check)
	fmt.Fprint(writer,check[0])
}