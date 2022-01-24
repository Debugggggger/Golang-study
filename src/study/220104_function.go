package main

import "fmt"

func main(){
	line_by_line("슈비루밥","슈비루바","뚜뚜")
	count,sum := multi(6,5,2,14,5,123,4)

	fmt.Println("갯수 : ",count, " 합계 : ",sum)
	avg := func(n ... int) int{
		s:=0
		for _,i := range n{
			s+=i
		}
		return s
	}
	result := avg(43,5,6,7,4)
	println(result);
}
func line_by_line (line...string) {
	for i,l := range line{
		fmt.Println(i,"번째 라인 ",l);
	}
	for j :=1 ; j<len(line);j++{
		fmt.Println(line[j],"이전 라인 ",line[j-1]);
	}
}
func multi(n...int)(cnt int, sum int){
	for _,e := range n {
		cnt++;
		sum+=e;
	}
	return cnt,sum
}