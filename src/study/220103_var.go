package main

import "fmt"

func main(){
	var a int = 0
	b:= 3
	var f float32 = 11.
	println(a);
	println(b);
	println(f);
	fmt.Println("Hello World")

	if(a==0){
		println("true")
	}

	sum := func(n ...int) int{
		s:=0
		for _, i := range n{
			s+=2*i
		}
		return s
	}
	result := sum(1,2,3,4,5)
	println(result)

	line_by_line("1","10","100","1000","10000")
	println(multi(3,4,5,6,7,8,9));
	cnt,sum2 := multi(3,4,5,6,7,8,9)
	println(cnt,sum2)
}

func line_by_line (n ... string){
	for i,e := range n{
		println(i,"번째 line",e);
	}
}

func multi (n...int)(cnt int,sum int){
	for _,e := range n{
		cnt++;
		sum+=e;
	}
	return cnt,sum;
}