package main

import "fmt"

func main() {
	kingdom := &FirstKingdom{}
	k := &King{}
	k.Build("John Snow", kingdom)

	sl := []Kingdom{&SecondKingdom{}, &ThirdKingdom{}}
	s := &Level2{}
	start := &Level1{}
	start.NextLevel(s)
	if start.LevelUp(k, sl) {
		fmt.Println("Congratulations you are winner")
	} else {
		fmt.Println("Game over. You are dead")
	}
}
