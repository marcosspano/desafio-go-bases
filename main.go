package main

import (
	"fmt"
)

// "github.com/bootcamp-go/desafio-go-bases/internal/tickets"

func main() {

	messageTotalTickets, _ := GetTotalTickets("China")
	fmt.Println(messageTotalTickets)

	messageTotalPeriodo, _ := GetCountByPeriod("tarde")
	fmt.Println(messageTotalPeriodo)
}
