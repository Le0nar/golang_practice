package main

import (
	"fmt"
)

type Message struct {
	From string
	To   string
	Content string
}

func main() {
	ch := make(chan Message)

	for i := 0; i < 5; i++ {
		go sendMessage(ch, i)
	}

	messageList := make([]Message, 5)

	for i := range messageList {
		mes := <-ch
		messageList[i] = mes
	}

	for _, v := range messageList {
		fmt.Printf("%v\n", v)
	}
}

func sendMessage(ch chan Message, index int) {
	content := fmt.Sprintf("Hello, index: %d", index)
	message := Message{From: "Ya", To: "Ka", Content: content}

	ch <- message
}
