package main

import (
	"fmt"
	"time"
)

type Coffee struct {
	kind string // kind of coffee
	name string // name of the customer
	id   int    // id to show it is executed async
}

func order(coffee Coffee, coffeChannel chan Coffee) {
	switch coffee.kind {
	case "cappuccino":
		time.Sleep(1000 * time.Millisecond)
		coffeChannel <- coffee
		break
	case "espresso":
		time.Sleep(2000 * time.Millisecond)
		coffeChannel <- coffee
		break
	case "flat white":
		time.Sleep(3000 * time.Millisecond)
		coffeChannel <- coffee
		break
	case "macchiato":
		time.Sleep(4000 * time.Millisecond)
		coffeChannel <- coffee
		break
	}
}

func main() {
	coffeChannel := make(chan Coffee)

	go order(Coffee{"macchiato", "Roman", 1}, coffeChannel)
	go order(Coffee{"espresso", "Marcel", 2}, coffeChannel)
	go order(Coffee{"flat white", "Carsten", 3}, coffeChannel)
	go order(Coffee{"cappuccino", "Max", 4}, coffeChannel)

	for i := 0; i < 4; i++ {
		coffee := <-coffeChannel
		fmt.Println(fmt.Sprintf("%d: Coffee for %s ready", coffee.id, coffee.name))
	}
}
