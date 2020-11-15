package main

import "fmt"

func main() {
	/*facade := newGameFacade()
	facade.createCharacter()
	facade.startGame(facade.createLevels())*/
	db := GetSingletonDatabase()
	number := db.GetNumber("Kuka")
	fmt.Println("Phone number of Kuka is", number)
}


