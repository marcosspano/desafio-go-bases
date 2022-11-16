package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id          int
	Nome        string
	Email       string
	PaisDestino string
	HoraVoo     string
	Preco       float64
}

// Quantidade de pessoas que viajam para um determinado país
func GetTotalTickets(destination string) (string, error) {

	f, _ := os.Open("tickets.csv")
	r := csv.NewReader(f)

	count := 0

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(record[0])
		preco, _ := strconv.ParseFloat(record[5], 8)
		ticket := Ticket{
			Id:          id,
			Nome:        record[1],
			Email:       record[2],
			PaisDestino: record[3],
			HoraVoo:     record[4],
			Preco:       preco,
		}

		if ticket.PaisDestino == destination {
			count++
		}

	}

	message := fmt.Sprintf("Total de viagens para o país %s: %d", destination, count)

	return message, nil
}

// Quantas pessoas viajam em um determinado período
var ErrIntervalo = errors.New("Insira um intervalo válido. 00:00 às 06:00 ou madrugada | 07:00 às 12:00 ou manhã | 13:00 às 19:00 ou tarde | 20:00 às 23:00 ou noite")

func GetCountByPeriod(time string) (string, error) {

	f, _ := os.Open("tickets.csv")
	r := csv.NewReader(f)

	count := 0

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(record[0])
		preco, _ := strconv.ParseFloat(record[5], 8)
		ticket := Ticket{
			Id:          id,
			Nome:        record[1],
			Email:       record[2],
			PaisDestino: record[3],
			HoraVoo:     record[4],
			Preco:       preco,
		}

		periodo, err := strconv.Atoi(strings.Split(ticket.HoraVoo, ":")[0])
		if err != nil {
			panic(err)
		}

		switch time {
		case "madrugada":
			if periodo >= 00 && periodo <= 06 {
				count++
			}
		case "manhã":
			if periodo >= 07 && periodo <= 12 {
				count++
			}
		case "tarde":
			if periodo >= 13 && periodo <= 19 {
				count++
			}
		case "noite":
			if periodo >= 20 && periodo <= 23 {
				count++
			}
			// default:
			// 	fmt.Errorf("Intervalo inválido: %w", ErrIntervalo)

		}
	}

	message := fmt.Sprintf("Total de pessoas que viajam no perído da %s: %d", time, count)
	return message, nil

}

// Média de pessoas que viajam para um determinado país
func AverageDestination() (string, error) {

	paises := make(map[string]int)

	f, _ := os.Open("tickets.csv")
	r := csv.NewReader(f)

	count := 0

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(record[0])
		preco, _ := strconv.ParseFloat(record[5], 8)
		ticket := Ticket{
			Id:          id,
			Nome:        record[1],
			Email:       record[2],
			PaisDestino: record[3],
			HoraVoo:     record[4],
			Preco:       preco,
		}

		if paises[ticket.PaisDestino] == 0 {
			paises[ticket.PaisDestino]++
		} else {
			paises[ticket.PaisDestino]++
		}

		count++

	}

	media := count / len(paises)

	message := fmt.Sprintf("Média de viagens: %d", media)

	return message, nil
}
