package main

import "fmt"

type Subject interface {
	register(Observer)
	unregister(Observer)
	notifyAll()
}

type Rating struct {
	subscribers []Observer
	list        []string
}

func (c *Rating) register(observer Observer) {
	c.subscribers = append(c.subscribers, observer)
}

func (c *Rating) unregister(observer Observer) {
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

func (c *Rating) notifyAll() {
	for _, observer := range c.subscribers {
		observer.update(c.list)
	}
}

func (c *Rating) addNewChar(name string) {
	c.list = append(c.list, name)
	c.notifyAll()
}

func (c *Rating) removeChar(name string) {
	for i, nameOfmovie := range c.list {
		if nameOfmovie == name {
			c.list = remove(c.list, i)
		}
	}
	c.notifyAll()
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

type Observer interface {
	update([]string)
}

func (p *King) update(list []string) {
	fmt.Println("Hello,", p.GetCharacter().Name)
	fmt.Println("We have new character in  game:")
	for i, p := range list {
		fmt.Println(i, p)
	}
}

func (p *Queen) update(list []string) {
	fmt.Println("Hello,", p.GetCharacter().Name)
	fmt.Println("We have new character in the game:")
	for i, p := range list {
		fmt.Println(i, p)
	}
}
