package main

import (
	"fmt"
)
type testCase struct{
	distance int
	r1 int
	r2 int
}
func main(){
	tc := []testCase{}
	var n int
	
	fmt.Scanln(&n)
	for i:=0 ; i<n ; i++{
		var x1,y1,r1,x2,y2,r2 int
		fmt.Scanln(&x1,&y1,&r1,&x2,&y2,&r2)
		tc = append(tc,testCase{(x1-x2)*(x1-x2)+(y1-y2)*(y1-y2),r1,r2})
	}
	for i:=0;i<len(tc);i++{
		if tc[i].r1==tc[i].r2&& tc[i].distance==0{
			fmt.Println(-1)
		}else{
			switch  {
			case tc[i].distance > (tc[i].r1+tc[i].r2)*(tc[i].r1+tc[i].r2) :
				fmt.Println(0)
			case tc[i].distance == (tc[i].r1+tc[i].r2)*(tc[i].r1+tc[i].r2) :
				fmt.Println(1)
			case tc[i].distance < (tc[i].r1+tc[i].r2)*(tc[i].r1+tc[i].r2) :
				switch {
					case tc[i].distance > (tc[i].r1-tc[i].r2)*(tc[i].r1-tc[i].r2) :
						fmt.Println(2)
					case tc[i].distance == (tc[i].r1-tc[i].r2)*(tc[i].r1-tc[i].r2) :
						fmt.Println(1)
					case tc[i].distance < (tc[i].r1-tc[i].r2)*(tc[i].r1-tc[i].r2) :
						fmt.Println(0)
				}
			}
		}
	}
}
