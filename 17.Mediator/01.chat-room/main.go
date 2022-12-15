package main

import "fmt"

type Person struct {
	Name    string
	Room    *ChatRoom
	chatlog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: '%s'", sender, message)
	p.chatlog = append(p.chatlog, s)
	fmt.Printf("[%s's chat session] %s\n", p.Name, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
	people []*Person
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

func (c *ChatRoom) Message(source, destination, message string) {
	for _, p := range c.people {
		if p.Name == destination {
			p.Receive(source, message)
		}
	}
}

func main() {
	room := NewChatRoom()

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "glad you could join us!")
}

// Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects.
// The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.
