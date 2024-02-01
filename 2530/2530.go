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

	var H, M, S, time int
	fmt.Fscanf(reader, "%d %d %d\n", &H, &M, &S)
	fmt.Fscanf(reader, "%d\n", &time)

	S += time%60
	time -= (time%60)
	M += ((time/60)%60)
	H += (time/3600)

	if S >= 60 {
		M += (S/60)
		S %= 60
	}
	if M >= 60 {
		H += (M/60)
		M %= 60
	}
	H %= 24

	fmt.Fprintf(writer, "%d %d %d\n", H, M, S)
}