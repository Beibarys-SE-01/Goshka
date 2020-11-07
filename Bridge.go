package main

import "fmt"

//First hierarchy
type Kingdom interface {
	GetKingdom() string
}

type FirstKingdom struct{}

func (n *FirstKingdom) GetKingdom() string {
	return "The North"
}

type SecondKingdom struct{}

func (n *SecondKingdom) GetKingdom() string {
	return "The Kingdom of Mountain and Vale"
}

type ThirdKingdom struct{}

func (n *ThirdKingdom) GetKingdom() string {
	return "Kingdom of the Isles and Rivers"
}

type FourthKingdom struct{}

func (n *FourthKingdom) GetKingdom() string {
	return "The Kingdom of the Rock"
}

type FifthKingdom struct{}

func (n *FifthKingdom) GetKingdom() string {
	return "The Kingdom of the Stormlands"
}

type SixthKingdom struct{}

func (n *SixthKingdom) GetKingdom() string {
	return "The Kingdom of the Reach"
}

type SeventhKingdom struct{}

func (n *SeventhKingdom) GetKingdom() string {
	return "The Principality of Dorne"
}

//Second hierarchy
type iCharacter interface {
	Attack()
	Protect()
	Upgrade()
	GetCharacter() *Character
}

type King struct {
	c *Character
}


func (k *King) GetCharacter() *Character {
	return k.c
}

func (k *King) Upgrade() {
	k.GetCharacter().Power += 50
	k.GetCharacter().Dexterity += 20
	k.GetCharacter().Protection += 30
	k.GetCharacter().Reaction += 10
	fmt.Printf("Now You are King. Your characteristics:\nPower: %v\nDexterity: %v\nProtection: %v\nReaction: %v\n",
		k.GetCharacter().Power, k.GetCharacter().Dexterity, k.GetCharacter().Protection, k.GetCharacter().Reaction)
}

func (k *King) Protect() {
	fmt.Println("You will regret that you were born in this world - Said King and defend his Kingdom")
}

func (k *King) Attack() {
	fmt.Println("This land is gonna be mine - Said King and attack ")
}

type Queen struct {
	c *Character
}


func (k *Queen) GetCharacter() *Character {
	return k.c
}

func (k *Queen) Upgrade() {
	k.GetCharacter().Power += 30
	k.GetCharacter().Dexterity += 30
	k.GetCharacter().Protection += 10
	k.GetCharacter().Reaction += 50
	fmt.Printf("Now You are Queen. Your characteristics:\nPower: %v\nDexterity: %v\nProtection: %v\nReaction: %v\n",
		k.GetCharacter().Power, k.GetCharacter().Dexterity, k.GetCharacter().Protection, k.GetCharacter().Reaction)
}

func (k *Queen) Protect() {
	fmt.Println("GIVE ME MY ARMOR, I'm going to protect my Kingdom - Said Queen and defend her Kingdom")
}

func (k *Queen) Attack() {
	fmt.Printf("Let's make %v Great again - Said Queen and attack\n", k.c.Kingdom.GetKingdom())
}

type Knight struct {
	c *Character
}

func (k *Knight) GetCharacter() *Character {
	return k.c
}

func (k *Knight) Upgrade() {
	k.GetCharacter().Power += 0
	k.GetCharacter().Dexterity += 40
	k.GetCharacter().Protection += 50
	k.GetCharacter().Reaction += 40
	fmt.Printf("Now You are Knight. Your characteristics:\nPower: %v\nDexterity: %v\nProtection: %v\nReaction: %v\n",
		k.GetCharacter().Power, k.GetCharacter().Dexterity, k.GetCharacter().Protection, k.GetCharacter().Reaction)
}

func (k *Knight) Protect() {
	fmt.Printf("God save %v - said Knight and protect his Homeland\n", k.c.Kingdom.GetKingdom())
}

func (k *Knight) Attack() {
	fmt.Printf("God bless %v - said Knight and attack\n", k.c.Kingdom.GetKingdom())
}

type HandOfKing struct {
	c *Character
}

func (k *HandOfKing) GetCharacter() *Character {
	return k.c
}

func (k *HandOfKing) Upgrade() {
	k.GetCharacter().Power += 20
	k.GetCharacter().Dexterity += 50
	k.GetCharacter().Protection += 20
	k.GetCharacter().Reaction += 30
	fmt.Printf("Now You are Hand Of the King. Your characteristics:\nPower: %v\nDexterity: %v\nProtection: %v\nReaction: %v\n",
		k.GetCharacter().Power, k.GetCharacter().Dexterity, k.GetCharacter().Protection, k.GetCharacter().Reaction)
}

func (k *HandOfKing) Protect() {
	fmt.Println("This strategy is the best for protection. Let's go my friends - Give his advice and save his Homeland")
}

func (k *HandOfKing) Attack() {
	fmt.Println("Now we have enough power to attack other Kingdom - said Hand of the King and conquered a neighboring kingdom")
}
