package main

import (
	"fmt"
	"github.com/Julienyla/Elvekryssing/event"
)

func main() {
	fmt.Println(event.ViewEvent())

	event.AllGoingOver()
	fmt.Println(event.ViewEvent())

	event.AllOver()
	fmt.Println(event.ViewEvent())


}