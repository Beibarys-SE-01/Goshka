package main

import (
	"fmt"
	"time"
)

type Level interface {
	LevelUp(iCharacter, []Kingdom) bool
	NextLevel(Level)
}

type Level1 struct {
	Level2 Level
}

func (l *Level1) LevelUp(c iCharacter, k []Kingdom) bool {
	startTime := time.Now()
	fmt.Println("Level 1")
	fmt.Println("загадка 1")
	var x string
	fmt.Scan(&x)
	duration := time.Since(startTime)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your dexterity\n", k[0].GetKingdom())
		c.GetCharacter().Dexterity -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		if duration.Seconds() > 10 {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return l.Level2.LevelUp(&CharWithKnife{char: c}, k)
}

func (l *Level1) NextLevel(level Level) {
	l.Level2 = level
}

type Level2 struct {
	Level3 Level
}

func (l *Level2) LevelUp(c iCharacter, k []Kingdom) bool {
	startTime := time.Now()
	fmt.Println("Level 2")
	fmt.Println("загадка 2")
	var x string
	fmt.Scan(&x)
	duration := time.Since(startTime)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your protection\n", k[0].GetKingdom())
		c.GetCharacter().Protection -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		if duration.Seconds() > 10 {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return true //l.Level3.LevelUp(&CharWithShield{char: c}, k)
}

func (l *Level2) NextLevel(level Level) {
	l.Level3 = level
}

/*type Level3 struct {
	Level4 Level
}

func (l *Level3) LevelUp(c iCharacter, k []Kingdom) bool {
	fmt.Println("Level 3")
	fmt.Println("загадка 3")
	var x string
	fmt.Scan(&x)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your reaction\n", k[0].GetKingdom())
		c.GetCharacter().Reaction -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return l.Level4.LevelUp(&CharWithPotion{char: c}, k)
}

func (l *Level3) NextLevel(level Level) {
	l.Level4 = level
}

type Level4 struct {
	Level5 Level
}

func (l *Level4) LevelUp(c iCharacter, k []Kingdom) bool {
	fmt.Println("Level 4")
	fmt.Println("загадка 4")
	var x string
	fmt.Scan(&x)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your power\n", k[0].GetKingdom())
		c.GetCharacter().Power -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return l.Level5.LevelUp(&CharWithDragon{char: c}, k)
}

func (l *Level4) NextLevel(level Level) {
	l.Level5 = level
}

type Level5 struct {
	Level6 Level
}

func (l *Level5) LevelUp(c iCharacter, k []Kingdom) bool {
	fmt.Println("Level 1")
	fmt.Println("загадка 1")
	var x string
	fmt.Scan(&x)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your dexterity\n", k[0].GetKingdom())
		c.GetCharacter().Dexterity -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return l.Level6.LevelUp(&CharWithKnife{char: c}, k)
}

func (l *Level5) NextLevel(level Level) {
	l.Level6 = level
}

type Level6 struct {
	Level7 Level
}

func (l *Level6) LevelUp(c iCharacter, k []Kingdom) bool {
	fmt.Println("Level 1")
	fmt.Println("загадка 1")
	var x string
	fmt.Scan(&x)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your dexterity\n", k[0].GetKingdom())
		c.GetCharacter().Dexterity -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return l.Level7.LevelUp(&CharWithKnife{char: c}, k)
}

func (l *Level6) NextLevel(level Level) {
	l.Level7 = level
}

type Level7 struct {
	Level2 Level
}

func (l *Level7) LevelUp(c iCharacter, k []Kingdom) bool {
	fmt.Println("Level 1")
	fmt.Println("загадка 1")
	var x string
	fmt.Scan(&x)
	for !CheckAnswer(x, "correctAnswer") {
		fmt.Printf("Oops your answer is wrong. %v has been attacked your Kingdom.\n -20points from your dexterity\n", k[0].GetKingdom())
		c.GetCharacter().Dexterity -= 20
		if !c.GetCharacter().CheckSkills() {
			return false
		}
		fmt.Scan(&x)
	}
	//c.Attack()
	return true
}

func (l *Level7) NextLevel(level Level) {

}*/

func CheckAnswer(string1 string, string2 string) bool {
	if string2 == string1 {
		return true
	}
	return false
}
