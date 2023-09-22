package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	check := make([]int, n+1)
    check[0] = 0
	for i:=1 ; i<=n ; i++{
		check[i] = check[i-1]+1
		for j:=1 ; j*j<=i ; j++ {
			check[i] = int(math.Min(float64(check[i]), float64(check[i-j*j]+1)))
		}
	}
	fmt.Fprintln(writer,check[n])
}