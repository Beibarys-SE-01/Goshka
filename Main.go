package main

func main() {
	facade := newGameFacade()
	facade.createCharacter()
	facade.startGame(facade.createLevels())
}


