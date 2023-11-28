package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Read the entire file
	readTicketFile()
	
	// Read file by line
	readTicketFileByline()
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

func readTicketFileByline() {
	file, err := os.Open("../createfile/tickets/Petr Petrov_0.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        fmt.Println(fileScanner.Text())
    }
}
