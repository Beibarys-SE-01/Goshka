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
func (c *CharWithKnife) Kingdom() Kingdom {
	return c.char.Kingdom()
}
func (c *CharWithKnife) Protect() {
	c.char.Protect()
}

func (c *CharWithKnife) Attack() {
	c.char.Attack()
	fmt.Println("You are attacking with knife")
}
type CharWithShield struct {
	char iCharacter
}

func (c *CharWithShield) Upgrade() {
	c.char.GetCharacter().Protection += 15
	fmt.Println("Now you have Shield in your arsenal. Your protection is ", c.char.GetCharacter().Protection)
}
func (c *CharWithShield) Kingdom() Kingdom {
	return c.char.Kingdom()
}
func (c *CharWithShield) GetCharacter() *Character {
	return c.char.GetCharacter()
}

func (c *CharWithShield) Protect() {
	c.char.Protect()
	fmt.Println("You are protecting with shield")
}

func (c *CharWithShield) Attack() {
	c.char.Attack()
}

type CharWithPotion struct {
	char iCharacter
}

func (c *CharWithPotion) Upgrade() {
	c.char.GetCharacter().Reaction += 20
	fmt.Println("Now you have Potion in your arsenal. Your reaction is ", c.char.GetCharacter().Reaction)
}
func (c *CharWithPotion) Kingdom() Kingdom {
	return c.char.Kingdom()
}
func (c *CharWithPotion) GetCharacter() *Character {
	return c.char.GetCharacter()
}

func (c *CharWithPotion) Protect() {
	c.char.Protect()
}

func (c *CharWithPotion) Attack() {
	c.char.Attack()
	fmt.Println("Wow you have a good reaction")
}
type CharWithDragon struct {
	char iCharacter
}

func (c *CharWithDragon) Upgrade() {
	c.char.GetCharacter().Power += 50
	fmt.Println("Now you have Dragon in your arsenal. Your power is ", c.char.GetCharacter().Power)
}
func (c *CharWithDragon) Kingdom() Kingdom {
	return c.char.Kingdom()
}
func (c *CharWithDragon) GetCharacter() *Character {
	return c.char.GetCharacter()
}

func (c *CharWithDragon) Protect() {
	c.char.Protect()
	fmt.Println("Your Dragon save your life")
}

func (c *CharWithDragon) Attack() {
	c.char.Attack()
}

type BecomeNobody struct {
	char iCharacter
}

func (c *BecomeNobody) GetCharacter() *Character {
	return c.char.GetCharacter()
}
func (c *BecomeNobody) Kingdom() Kingdom {
	return c.char.Kingdom()
}
func (c *BecomeNobody) Upgrade() {
	c.char.GetCharacter().Power += 50
	c.char.GetCharacter().Dexterity += 50
	c.char.GetCharacter().Reaction += 50
	c.char.GetCharacter().Protection += 50
	fmt.Println("You became nobody:)")
}
func (c *BecomeNobody) Protect() {
	fmt.Println("Your skills help you to pretend White Walker")
}

func (c *BecomeNobody) Attack() {
	fmt.Println("Wow you kill White Walker")
}