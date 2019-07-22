package main

import (
	"fmt"

	"gomoddemo/messages"
)

type Sayer interface {
	Say() string
}

func main() {
	message := messages.NewHello()
	fmt.Println(message.Say())
}

