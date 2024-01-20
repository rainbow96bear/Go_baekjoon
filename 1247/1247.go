package main

import (
	"bufio"
	"fmt"
	"os"
)
type Num struct {
	big int64
	small int64
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	count := 3
	divide := 100000
	for i:= 0 ; i<count ; i++ {
		var N int
		fmt.Fscanf(reader, "%d\n", &N)
		var total Num
		for n:=0 ; n<N ; n++ {
			var input int64
			fmt.Fscanf(reader, "%d\n", &input)
			num := Num{input/int64(divide), input%int64(divide)}
			total = Num{total.big+num.big, total.small+num.small}
		}
		if total.big > 0 || (total.big == 0 && total.small>0) {
			fmt.Fprintln(writer, "+")
		}else if total.big == 0 && total.small == 0{
			fmt.Fprintln(writer, "0")
		}else {
			fmt.Fprintln(writer, "-")
		}
	}
}