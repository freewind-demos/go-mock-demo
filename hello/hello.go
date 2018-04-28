package hello

import "fmt"

type Talker interface {
	SayHello(word string) (response string)
}

type Company struct {
	Usher Talker
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		name: "Freewind",
	}
}

func (p *Person) SayHello(name string) (word string) {
	return fmt.Sprintf("Hello, %s, welcome to our company, my name is %s", name, p.name)
}

func NewCompany(talker Talker) *Company {
	return &Company{
		Usher: talker,
	}
}

func (c *Company) Meeting(guestName string) string {
	return c.Usher.SayHello(guestName)
}
