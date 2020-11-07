package main

import "fmt"

type Character struct {
	Name                                   string
	Kingdom                                Kingdom
	Power, Dexterity, Reaction, Protection int
}

type CharBuilder struct {
	c *Character
}

func (b *CharBuilder) Build() *Character {
	return b.c
}

func (b *CharBuilder) Name(name string) *CharBuilder {
	b.c.Name = name
	return b
}

func (b *CharBuilder) LiveIn(k Kingdom) *CharBuilder {
	b.c.Kingdom = k
	fmt.Println(b.c.Kingdom.GetKingdom())
	return b
}

func (b *CharBuilder) WithPower(power int) *CharBuilder {
	b.c.Power = power
	return b
}

func (b *CharBuilder) WithDexterity(d int) *CharBuilder {
	b.c.Dexterity = d
	return b
}

func (b *CharBuilder) WithReaction(r int) *CharBuilder {
	b.c.Reaction = r
	return b
}

func (b *CharBuilder) WithProtection(p int) *CharBuilder {
	b.c.Protection = p
	return b
}

func (c *Character) CheckSkills() bool {
	if c.Reaction*c.Power*c.Dexterity*c.Protection <= 0 {
		return false
	}
	return true
}

func (c *Character) String() string {
	return fmt.Sprintf("%v with Power: %v, With Dexterity %v, With Reaction %v, With Protection %v",
		c.Name, c.Power, c.Dexterity, c.Reaction, c.Protection)
}
func Build(name string, kingdom Kingdom) *Character{
	c := NewCharBuilder()
	c.Name(name).
		LiveIn(kingdom).
		WithPower(50).
		WithDexterity(50).
		WithProtection(50).
		WithReaction(50)
	return c.Build()
}


func NewCharBuilder() *CharBuilder {
	return &CharBuilder{&Character{}}
}
