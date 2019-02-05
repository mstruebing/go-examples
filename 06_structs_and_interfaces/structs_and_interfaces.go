package main

import (
	"fmt"
)

// declare a struct named Person
type Person struct {
	name string
	age  int
}

func (p Person) sayName() {
	fmt.Println("My name is", p.name)
}

// function to be called on the Person struct
func (p Person) sayAge() {
	fmt.Printf("I'm %d years old\n", p.age)
}

// function to be called on the person struct when changing
// data (pointer/asterisk/call by reference)
func (p *Person) birthday() {
	p.age++
}

// the change on the person would never reach the outside world
// call by value, not by reference
func (p Person) birthday2() {
	p.age += 1
}

// Use Person struct inside of Pupil
// duck typing
type Pupil struct {
	Person
	grade int
}

// interface has one method of type Say(string)
type Speaker interface {
	Say(string)
}

// interface implementation of Say method to become a Speaker
func (p Person) Say(msg string) {
	fmt.Println(p.name + ": " + msg)
}

// speakSomething gets a Speaker (a struct which implements the Speaker interface)
func speakSomething(s Speaker) {
	s.Say("something")
}

// struct of type Dog
type Dog struct {
	name string
}

// interface implementation of Say method to become a Speaker
func (d Dog) Say(msg string) {
	fmt.Println(d.name + ": bark bark bark")
}

func main() {
	// create an instance of the struct
	me := Person{"Max", 28}
	fmt.Println(me) // My name is Max

	me.sayAge()
	// change value and call again
	me.birthday()
	me.sayAge()

	// directly manipulate struct
	me.age = 10
	me.sayAge()

	// reset
	me = Person{"Max", 28}
	me.birthday2()
	me.sayAge()

	fmt.Println("--------")

	// create a Pupil and use duck typing
	schoolkid := Pupil{Person{"Max", 14}, 7}
	schoolkid.sayAge()
	schoolkid.birthday()
	schoolkid.sayAge()

	fmt.Println("--------")

	// use the Say method
	me.Say("Hello")

	// use the speakSomething func which gets any Speaker
	speakSomething(me)
	speakSomething(schoolkid)

	// we define that a Dog is also a Speaker
	dog := Dog{"Bello"}
	speakSomething(dog)
}
