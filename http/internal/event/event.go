package event

type Event struct {
	Id      string `json:"id"`
	UserId  int    `json:"userId" `
	Date    string `json:"date"`
	Content string `json:"content"`
}

type EventDto struct {
	Id      string `json:"id"`
	UserId  int    `json:"userId" `
	Date    string `json:"date"`
	Content string `json:"content"`
}
