package event

type Event struct {
	Id      string `json:"id"`
	UserId  int    `json:"userId" `
	Date    string `json:"date"`
	Content string `json:"content"`
}

type CreateEventDto struct {
	UserId  int    `json:"userId" binding:"required"`
	Date    string `json:"date" binding:"required"`
	Content string `json:"content" binding:"required"`
}
