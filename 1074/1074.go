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

	var N, r, c int
	fmt.Fscan(reader, &N, &r, &c)

	answer := 0 
	for i:=1 ; i<=N ; i++ {
		nowCheck := int(math.Pow(2.0, float64(N-i)))

		if r < nowCheck {
			if nowCheck <= c {
				answer += nowCheck*nowCheck
				c -= nowCheck
			}
		}else {
			r -= nowCheck 
			if c < nowCheck {
				answer += nowCheck * nowCheck * 2
			}else {
				answer += nowCheck * nowCheck * 3
				c -= nowCheck
			}
		}
	}
	fmt.Fprint(writer, answer)
}