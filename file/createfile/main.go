package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	ticket := Ticket{
		Id: 0,
		DepartureTime: "2023-06-22 19:10",
		DepartureCity: "Moscow",
		ArrivalTime: "2023-06-22 21:50",
		ArrivalCity: "Izhevsk",
		FlightDuration: "1h 30m",
		Passanger: "Petr Petrov",
		BaggageKg: 20,
		PriceRub: 3500,
	}

	createTicketFile(ticket)
}

type Ticket struct {
	Id int
	DepartureTime string
	DepartureCity string
	ArrivalTime string
	ArrivalCity string
	FlightDuration string
	Passanger string
	BaggageKg float64
	PriceRub float64
}

func createTicketFile(ticket Ticket) {
	// TODO: move to directroy

	err := os.MkdirAll("tickets", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	ticketContent := ""

	ticketContent += fmt.Sprintf("Departure: %s %s \n", ticket.DepartureCity, ticket.DepartureTime)
	ticketContent += fmt.Sprintf("Arrival: %s %s \n", ticket.ArrivalCity, ticket.ArrivalTime)
	ticketContent += fmt.Sprintf("Flight Time: %s \n", ticket.FlightDuration)
	ticketContent += fmt.Sprintf("Passanger: %s \n", ticket.Passanger)
	changedBagage := strconv.FormatFloat(ticket.BaggageKg, 'f', -1, 64)
	ticketContent += fmt.Sprintf("Baggage: %sKg \n", changedBagage)
	changedPrice := strconv.FormatFloat(ticket.PriceRub, 'f', 2, 64)
	ticketContent += fmt.Sprintf("Price: %s Rub \n", changedPrice)

	changedTicketContent := []byte(ticketContent)

	ticketFileName := fmt.Sprintf("tickets/%s_%d.txt", ticket.Passanger, ticket.Id)

	os.WriteFile(ticketFileName, changedTicketContent, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
