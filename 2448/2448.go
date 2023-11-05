package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N, _ := strconv.Atoi(Input(scanner))
	
	Process(N, writer)
	
}

func Process(N int, writer *bufio.Writer) {
	zeroLine := "*"
	oneLine := "* *"
	twoLine := "*****"
	list := []string{zeroLine, oneLine, twoLine}
	fmt.Fprintln(writer, strings.Repeat(" ",N-1) + zeroLine + strings.Repeat(" ",N-1))
	fmt.Fprintln(writer, strings.Repeat(" ",N-2) + oneLine + strings.Repeat(" ",N-2))
	fmt.Fprintln(writer, strings.Repeat(" ",N-3) + twoLine + strings.Repeat(" ",N-3))
	for n:=3 ; n<N ;  {
		nextList := []string{}
		for i, v := range list {
			output := v+strings.Repeat(" ", 6*(n/3)-6*(i/3)-len(v))+v
			nextList = append(nextList, output)
			fmt.Fprintln(writer, strings.Repeat(" ",N-1-n)+output+strings.Repeat(" ",N-1-n))
			n++
		}
		if n == N {
			break
		}
		list = append(list, nextList...)
	}
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}