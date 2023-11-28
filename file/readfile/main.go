package main

import (
	"io"
	"log"
	"os"
)

func main() {
	readTicketFile()
}

func readTicketFile() {
	file, err := os.Open("../createfile/tickets/Petr Petrov_0.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	
	os.Stdout.Write(data)
}
