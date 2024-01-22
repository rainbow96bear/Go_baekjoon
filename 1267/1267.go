package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)
	for n:=0 ; n<N ; n++ {
		fmt.Fscanf(reader, "%d", &list[n])
	}
	fmt.Fscanf(reader, "\n")

	Yprice := 0
	Mprice := 0
	for i:=0 ; i<N ; i++ {
		num := list[i]
		for num >= 0 {
			Yprice += 10
			num -=30 
		}
	}
	for i:=0 ; i<N ; i++ {
		num := list[i]
		for num >= 0 {
			Mprice += 15
			num -=60 
		}
	}

	if Yprice > Mprice {
		fmt.Fprintf(writer, "M %d\n", Mprice)
	}else if Yprice == Mprice {
		fmt.Fprintf(writer, "Y M %d", Yprice)
	}else {
		fmt.Fprintf(writer, "Y %d", Yprice)
	}
}