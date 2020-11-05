package main

import "fmt"

type CharWithKnife struct {
	char iCharacter
}

func (c *CharWithKnife) Upgrade() {
	c.char.GetCharacter().Dexterity += 10
	fmt.Println("Now you have Knife in your arsenal. Your dexterity is ", c.char.GetCharacter().Dexterity)
}

func (c *CharWithKnife) GetCharacter() *Character {
	return c.char.GetCharacter()
}

type CharWithShield struct {
	char iCharacter
}

func (c *CharWithShield) Upgrade() {
	c.char.GetCharacter().Protection += 15
	fmt.Println("Now you have Shield in your arsenal. Your protection is ", c.char.GetCharacter().Protection)
}

func (c *CharWithShield) GetCharacter() *Character {
	return c.char.GetCharacter()
}

type CharWithPotion struct {
	char iCharacter
}

func (c *CharWithPotion) Upgrade() {
	c.char.GetCharacter().Reaction += 20
	fmt.Println("Now you have Potion in your arsenal. Your reaction is ", c.char.GetCharacter().Reaction)
}

func (c *CharWithPotion) GetCharacter() *Character {
	return c.char.GetCharacter()
}

type CharWithDragon struct {
	char iCharacter
}

func (c *CharWithDragon) Upgrade() {
	c.char.GetCharacter().Power += 50
	fmt.Println("Now you have Dragon in your arsenal. Your protection is ", c.char.GetCharacter().Power)
}

func (c *CharWithDragon) GetCharacter() *Character {
	return c.char.GetCharacter()
}

type CharWithHelp struct {
	char iCharacter
}

func (c *CharWithHelp) Upgrade() {
	fmt.Println("Здесь будет подсказка")
}

type CharWithTime struct {
	char iCharacter
}

func (c *CharWithTime) Upgrade() {
	fmt.Println("Здесь будет time")
}

type BecomeNobody struct {
	char iCharacter
}

func (c *BecomeNobody) Upgrade() {
	c.char.GetCharacter().Power += 50
	c.char.GetCharacter().Dexterity += 50
	c.char.GetCharacter().Reaction += 50
	c.char.GetCharacter().Protection += 50
	fmt.Println("You became nobody:)")
}
