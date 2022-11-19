package core

type ApiPostPublication struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	Mineature string `json:"mineature"`
	Topic     string `json:"topic"`
}

type ApiGetPublication struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	Title        string `json:"title"`
	Mineature    string `json:"mineature"`
	Author       string `json:"author"`
	Date         int    `json:"date"`
	Topic        string `json:"topic"`
	Introduction string `json:"introduction"`
}

type ApiInformation struct {
	Page     int
	To       int
	Quantity int                 `json:"quantity"`
	Posts    []ApiGetPublication `json:"publications"` // no va a guardar el contenido , solo una preview
	Topics   []string            `json:"topics"`
}
