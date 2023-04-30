package main

import(
	"time"
	"github.com/yoshiyoshiharu/typing_game/typing"
)

func countDown(ch chan <- string) {
	time.Sleep(10 * time.Second)
	ch <- "finished"
}

func question(ch chan <- int, point int) {
	var is_correct bool

	question := "aaa"
	println("question: ", question)

	is_correct = typing.Judge(question)

	if is_correct {
		point++
	}

	ch <- point

	println(point)
}

func main() {
	point := 0

	ch1 := make(chan string)
	ch2 := make(chan int)

	go countDown(ch1)
	go question(ch2, point)

L:
	for {
		select {
		case msg := <- ch1:
			println(msg)
			break L
		case <- ch2:
	}
}
