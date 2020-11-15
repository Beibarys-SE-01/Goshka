package main

import (
	"fmt"
	"strings"
)

type Level interface {
	LevelUp(iCharacter, []Kingdom) bool
	NextLevel(Level)
}

type Level1 struct {
	Level2 Level
}

func (l *Level1) LevelUp(c iCharacter, k []Kingdom) bool {
	c.Upgrade()
	fmt.Println("\nLevel 1")
	
	fmt.Printf("%v vs %v\n", k[0].GetKingdom(), c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("I am purple, yellow, red, and green\nThe King cannot reach me and neither can the Queen.\nI show my colors after the rain\nAnd only when the sun comes out again.")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "Rainbow") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your dexterity\n", k[0].GetKingdom())
		c.GetCharacter().Dexterity -= 20
		fmt.Println("Your dexterity: ", c.GetCharacter().Dexterity)
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return l.Level2.LevelUp(&CharWithKnife{char: c}, k)
}

func (l *Level1) NextLevel(level Level) {
	l.Level2 = level
}

type Level2 struct {
	Level3 Level
}

func (l *Level2) LevelUp(c iCharacter, k []Kingdom) bool {
	c.Upgrade()
	fmt.Println("\nLevel 2")
	fmt.Printf("%v vs %v\n", k[1].GetKingdom(), c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("What is full of holes but still holds water?")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "Sponge") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your protection\n", k[1].GetKingdom())
		c.GetCharacter().Protection -= 20
		fmt.Println("Your protection: ", c.GetCharacter().Protection)
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return l.Level3.LevelUp(&CharWithShield{char: c}, k)
}

func (l *Level2) NextLevel(level Level) {
	l.Level3 = level
}

type Level3 struct {
	Level4 Level
}

func (l *Level3) LevelUp(c iCharacter, k []Kingdom) bool {
	c.Upgrade()
	fmt.Println("\nLevel 3")
	fmt.Printf("%v vs %v\n", k[2].GetKingdom(), c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("What gets wet while drying?")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "Towel") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your reaction\n", k[2].GetKingdom())
		c.GetCharacter().Reaction -= 20
		fmt.Println("Your reaction: ", c.GetCharacter().Reaction)
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return l.Level4.LevelUp(&CharWithPotion{char: c}, k)
}

func (l *Level3) NextLevel(level Level) {
	l.Level4 = level
}

type Level4 struct {
	Level5 Level
}

func (l *Level4) LevelUp(c iCharacter, k []Kingdom) bool {
	c.Upgrade()
	fmt.Println("\nLevel 4")
	fmt.Printf("%v vs %v\n", k[3].GetKingdom(), c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("What gets bigger when more is taken away?")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "Hole") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your power\n", k[3].GetKingdom())
		c.GetCharacter().Power -= 20
		fmt.Println("Your power: ", c.GetCharacter().Power)
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return l.Level5.LevelUp(&CharWithDragon{char: c}, k)
}

func (l *Level4) NextLevel(level Level) {
	l.Level5 = level
}

type Level5 struct {
	Level6 Level
}

func (l *Level5) LevelUp(c iCharacter, k []Kingdom) bool {
	c.Upgrade()
	fmt.Println("\nLevel 5")
	fmt.Printf("%v vs %v\n", k[4].GetKingdom(), c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("If you’ve got me, you want to share me\nIf you share me, you haven’t kept me.\nWhat am I?")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "Secret") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -10 points from your dexterity and power\n", k[4].GetKingdom())
		c.GetCharacter().Dexterity -= 10
		c.GetCharacter().Power -= 10
		fmt.Println("Your dexterity: ", c.GetCharacter().Dexterity)
		fmt.Println("Your power: ", c.GetCharacter().Power)
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return l.Level6.LevelUp(c, k)
}

func (l *Level5) NextLevel(level Level) {
	l.Level6 = level
}

type Level6 struct {
	Level7 Level
}

func (l *Level6) LevelUp(c iCharacter, k []Kingdom) bool {
	fmt.Println("\nLevel 6")
	fmt.Printf("%v vs %v\n", k[5].GetKingdom(), c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("You have me today,\nTomorrow you'll have more;\nAs your time passes,\nI'm not easy to store;\nI don't take up space,\nBut I'm only in one place;\nI am what you saw,\nBut not what you see.\nWhat am I?")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "Memory") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -10 points from your reaction and protection\n", k[5].GetKingdom())
		c.GetCharacter().Reaction -= 10
		c.GetCharacter().Protection -= 10
		fmt.Println("Your reaction: ", c.GetCharacter().Reaction)
		fmt.Println("Your protection: ", c.GetCharacter().Protection)
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return l.Level7.LevelUp(&BecomeNobody{char: c}, k)
}

func (l *Level6) NextLevel(level Level) {
	l.Level7 = level
}

type Level7 struct {
}

func (l *Level7) LevelUp(c iCharacter, k []Kingdom) bool {
	c.Upgrade()
	fmt.Println("\nLevel 7")
	fmt.Printf("White Walker vs %v\n", c.GetCharacter().Kingdom.GetKingdom())
	a := choose()
	fmt.Println("There are 5 houses in five different colors. In each house lives a person with a different nationality. These five owners drink a certain type of beverage, smoke a certain brand of cigar and keep a certain pet. No owners have the same pet, smoke the same brand of cigar or drink the same beverage.")
	fmt.Println("The question is, who owns the fish?")
	fmt.Println("The Brit lives in the red house\nThe Swede keeps dogs as pets\nThe Dane drinks tea\nThe green house is on the left of the white house\nThe green house’s owner drinks coffee\nThe person who smokes Pall Mall rears birds\nThe owner of the yellow house smokes Dunhill\nThe man living in the center house drinks milk\nThe Norwegian lives in the first house\nThe man who smokes blends lives next to the one who keeps cats\nThe man who keeps horses lives next to the man who smokes Dunhill\nThe owner who smokes BlueMaster drinks beer\nThe German smokes Prince\nThe Norwegian lives next to the blue house\nThe man who smokes blend has a neighbor who drinks water")
	var x string
	fmt.Scan(&x)
	for !strings.EqualFold(x, "German") {
		fmt.Println("Oops your answer is wrong. White Walkers has been attacked your Kingdom.\n All your skills loose 20 points")
		c.GetCharacter().Dexterity -= 20
		c.GetCharacter().Protection -= 20
		c.GetCharacter().Power -= 20
		c.GetCharacter().Reaction -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	if a {
		c.Attack()
	}else{
		c.Protect()
	}
	return true
}

func (l *Level7) NextLevel(level Level) {

}

func choose() bool {
	fmt.Println("Do you want to attack or protect ?")
	var x string
	fmt.Scan(&x)
	if strings.EqualFold("attack", x){
		return true
	} else if strings.EqualFold("protect", x){
		return false
	}else {
		return choose()
	}
}