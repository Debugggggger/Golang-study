package main

func main() {
	println("==============start==============")
	chanel := make(chan string)

	go func() {
		chanel <- "english"
		chanel <- "korean"
		close(chanel)
	}()

	for msg := range chanel {
		transMSG := trans(msg)

		println(transMSG)
	}
	var a string
	a = <-chanel
	println(a)
	println("==============end==============")
}

func trans(ori string) (res string) {
	switch ori {
	case "english":
		res = "영어"
		break
	case "korean":
		res = "한국어"
		break
	}
	return res
}