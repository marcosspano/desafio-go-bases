package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	// "github.com/marcosspano/desafio-go-bases/tree/main/internal/tickets"
)

func main() {

	messageTotalTickets, _ := tickets.GetTotalTickets("China")
	fmt.Println(messageTotalTickets)

	messageTotalPeriodo, _ := tickets.GetCountByPeriod("tarde")
	fmt.Println(messageTotalPeriodo)
}
