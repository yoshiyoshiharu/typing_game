package main

import "github.com/yoshiyoshiharu/typing_game/typing"

func main() {
	var is_correct bool
	var point int;

	for {
		title := "aaa"
		println("title: ", title)

		is_correct = typing.Judge(title)

		if is_correct {
			point++
		}
		println("point: ", point)
	}

}
