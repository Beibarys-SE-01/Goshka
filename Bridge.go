package main

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
	//Attack()
	//Protect()
	Upgrade()
	GetCharacter() *Character
}

type King struct {
	c *Character
}

func (k *King) Build(name string, kingdom Kingdom) {
	c := NewCharBuilder()
	c.Name(name).
		LiveIn(kingdom).
		WithPower(100).
		WithDexterity(70).
		WithProtection(80).
		WithReaction(60)
	k.c = c.Build()
}

func (k *King) GetCharacter() *Character {
	return k.c
}

func (k *King) Upgrade() {}

type Queen struct {
	c *Character
}

func (k *Queen) Build(name string, kingdom Kingdom) {
	c := NewCharBuilder()
	c.Name(name).
		LiveIn(kingdom).
		WithPower(80).
		WithDexterity(80).
		WithProtection(60).
		WithReaction(100)
	k.c = c.Build()
}

func (k *Queen) GetCharacter() *Character {
	return k.c
}

func (k *Queen) Upgrade() {}

type Knight struct {
	c *Character
}

func (k *Knight) Build(name string, kingdom Kingdom) {
	c := NewCharBuilder()
	c.Name(name).
		LiveIn(kingdom).
		WithPower(50).
		WithDexterity(90).
		WithProtection(100).
		WithReaction(90)
	k.c = c.Build()
}

func (k *Knight) GetCharacter() *Character {
	return k.c
}

func (k *Knight) Upgrade() {}

type HandOfKing struct {
	c *Character
}

func (k *HandOfKing) Build(name string, kingdom Kingdom) {
	c := NewCharBuilder()
	c.Name(name).
		LiveIn(kingdom).
		WithPower(70).
		WithDexterity(100).
		WithProtection(70).
		WithReaction(80)
	k.c = c.Build()
}

func (k *HandOfKing) GetCharacter() *Character {
	return k.c
}

func (k *HandOfKing) Upgrade() {}
