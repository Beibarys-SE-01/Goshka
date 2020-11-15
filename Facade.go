package main

import "fmt"

type gameFacade struct {
	character iCharacter
	kingdoms []Kingdom
	listOfWinners listOfWinners
}

func newGameFacade() *gameFacade {
	facade := &gameFacade{
		character: nil,
		kingdoms: []Kingdom{&FirstKingdom{}, &SecondKingdom{}, &ThirdKingdom{}, &FourthKingdom{}, &FifthKingdom{}, &SixthKingdom{}, &SeventhKingdom{}},
		listOfWinners: listOfWinners{subscribers: []Observer{}},
	}
	return facade
}

/*func (f *gameFacade) winners()  {
	
}*/

func (f *gameFacade) characterName() string {
	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)
	return name
}

func (f *gameFacade) chooseKingdom() Kingdom {
	fmt.Println("Choose your kingdom:\n 1-The North\n 2-The Kingdom of Mountain and Vale\n" +
		" 3-Kingdom of the Isles and Rivers\n 4-The Kingdom of the Rock\n 5-The Kingdom of the Stormlands\n" +
		" 6-The Kingdom of the Reach\n 7-The Principality of Dorne")
	var choose int
	fmt.Scan(&choose)
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
	copy(f.kingdoms[choose-1:], f.kingdoms[choose:])
	f.kingdoms[len(f.kingdoms)-1] = nil
	f.kingdoms = f.kingdoms[:len(f.kingdoms)-1]
	return k
}

func (f *gameFacade) createCharacter() {
	ch := Build(f.characterName())
	k := f.chooseKingdom()
	fmt.Println("Choose your character:\n 1-King\n 2-Queen\n" +
		" 3-Knight\n 4-HandOfKing")

	var choose int
	fmt.Scan(&choose)
	var char iCharacter
	switch choose {
	case 1:
		char = &King{ ch, k}
	case 2:
		char = &Queen{ch, k}
	case 3:
		char = &Knight{ch, k}
	case 4:
		char = &HandOfKing{ch, k}
	}
	f.character = char
}

func (f *gameFacade) createLevels() Level {
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
	return l1
}

func (f *gameFacade) startGame(l1 Level) {

	f.listOfWinners.register(Build("John Snow"))
	f.listOfWinners.register(Build("Tomiris"))
	f.listOfWinners.register(Build("Zhansaya"))
	f.listOfWinners.register(Build("Beibarys"))
	if l1.LevelUp(f.character, f.kingdoms) {
		fmt.Println("\nCongratulations you are winner\n")
		f.listOfWinners.notifyAll(f.character)
	} else {
		fmt.Println("\nGame over. You are dead")
	}
}