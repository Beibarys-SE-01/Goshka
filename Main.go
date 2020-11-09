package main

import "fmt"

func main() {
	listOfWinners := listOfWinners{subscribers: []Observer{}}

	listOfWinners.register(Build("John Snow", &ThirdKingdom{}))
	listOfWinners.register(Build("Tomiris", &SecondKingdom{}))
	listOfWinners.register(Build("Zhansaya", &SeventhKingdom{}))
	listOfWinners.register(Build("Beibarys", &FirstKingdom{}))

	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)
	fmt.Println("Choose your kingdom:\n 1-The North\n 2-The Kingdom of Mountain and Vale\n" +
					" 3-Kingdom of the Isles and Rivers\n 4-The Kingdom of the Rock\n 5-The Kingdom of the Stormlands\n" +
					" 6-The Kingdom of the Reach\n 7-The Principality of Dorne")
	var choose int
	fmt.Scan(&choose)
	sl := []Kingdom{&FirstKingdom{}, &SecondKingdom{}, &ThirdKingdom{}, &FourthKingdom{}, &FifthKingdom{}, &SixthKingdom{}, &SeventhKingdom{}}
	var k Kingdom
	switch choose {
	case 1:
		k = &FirstKingdom{}
	case 2:
		k = &SecondKingdom{}
	case 3:
		k = &ThirdKingdom{}
	case 4:
		k = &FourthKingdom{}
	case 5:
		k = &FifthKingdom{}
	case 6:
		k = &SixthKingdom{}
	case 7:
		k = &SeventhKingdom{}
	}
	copy(sl[choose-1:], sl[choose:])
	sl[len(sl)-1] = nil
	sl = sl[:len(sl)-1]
	ch := Build(name, k)
	fmt.Println("Choose your character:\n 1-King\n 2-Queen\n" +
					" 3-Knight\n 4-HandOfKing")

	fmt.Scan(&choose)
	var char iCharacter
	switch choose {
	case 1:
		char = &King{ch}
	case 2:
		char = &Queen{ch}
	case 3:
		char = &Knight{ch}
	case 4:
		char = &HandOfKing{ch}
	}
	l7 := &Level7{}
	l6 := &Level6{}
	l5 := &Level5{}
	l4 := &Level4{}
	l3 := &Level3{}
	l2 := &Level2{}
	l1 := &Level1{}
	l6.NextLevel(l7)
	l5.NextLevel(l6)
	l4.NextLevel(l5)
	l3.NextLevel(l4)
	l2.NextLevel(l3)
	l1.NextLevel(l2)
	if l1.LevelUp(char, sl) {
		fmt.Println("\nCongratulations you are winner\n")
		listOfWinners.notifyAll(char)
	} else {
		fmt.Println("\nGame over. You are dead")
	}
}


