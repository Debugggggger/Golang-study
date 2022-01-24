package main

import (
	"io"
	"os"
)

func main(){
	fi, err := os.Open("C:\\Users\\antline\\Documents\\TGM\\golang\\study\\src\\1.txt")
	if( err != nil){
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("C:\\Users\\antline\\Documents\\TGM\\golang\\study\\src\\2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for{
		// 읽기
		cnt,err := fi.Read(buff)
		if err != nil && err != io.EOF{
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt ==0 {
			break
		}

		// 쓰기
		_,err = fo.Write(buff[:cnt])
		if err != nil{
			panic(err)
		}
	}

}