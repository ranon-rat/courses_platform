package core

type ApiPostPublication struct {
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	Mineature string   `json:"mineature"`
	Topics    []string `json:"topics"`
}

type ApiGetPublication struct {
	ID        string   `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	Mineature string   `json:"mineature"`
	Author    string   `json:"author"`
	Date      int      `json:"date"`
	Topics    []string `json:"topics"`
}

type ApiInformation struct {
	Quantity int                 `json:"quantity"`
	Posts    []ApiGetPublication `json:"publications"` // no va a guardar el contenido , solo una preview
	Topics   []string            `json:"topics"`
}
