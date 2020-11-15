package main

import "fmt"

type Subject interface {
	register(Observer)
	unregister(Observer)
	notifyAll()
}

type listOfWinners struct {
	subscribers []Observer
}

func (c *listOfWinners) register(observer Observer) {
	c.subscribers = append(c.subscribers, observer)
}

func (c *listOfWinners) unregister(observer Observer) {
	for i, sub := range c.subscribers {
		if sub == observer {
			c.subscribers = removeFromSlice(c.subscribers, i)
		}
	}
}

func removeFromSlice(s []Observer, i int) []Observer {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (c *listOfWinners) notifyAll(ch iCharacter) {
	for _, observer := range c.subscribers {
		observer.update(ch)
	}
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

type Observer interface {
	update(ch iCharacter)
}

func (p *Character) update(ch iCharacter) {
	fmt.Println("Hello,", p.Name)
	fmt.Printf("We have new character who win the game, %v from %v \n", ch.GetCharacter().Name, ch.Kingdom().GetKingdom())
}
