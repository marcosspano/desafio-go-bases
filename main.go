package main

import (
	"fmt"

	"github.com/marcosspano/desafio-go-bases/tree/main/internal/tickets/internal/tickets"
)

func main() {

	messageTotalTickets, _ := tickets.GetTotalTickets("China")
	fmt.Println(messageTotalTickets)

	messageTotalPeriodo, _ := tickets.GetCountByPeriod("tarde")
	fmt.Println(messageTotalPeriodo)
}
