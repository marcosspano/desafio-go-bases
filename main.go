package main

import (
	"fmt"

	"github.com/marcosspano/desafio-go-bases/tree/main/internal/tickets/internal/tickets"
)

type Ticket struct {
	Id          int
	Nome        string
	Email       string
	PaisDestino string
	HoraVoo     string
	Preco       float64
}

func main() {

	messageTotalTickets, _ := tickets.GetTotalTickets("China")
	fmt.Println(messageTotalTickets)

	messageTotalPeriodo, _ := tickets.GetCountByPeriod("tarde")
	fmt.Println(messageTotalPeriodo)

	tickets.AverageDestination()
}
