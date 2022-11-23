package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
)

const (
	createPostQuery = "INSERT INTO publications(title,content,mineature,author,topic,datePublication,introduction) VALUES(?1,?2,?3,?4,?5,?6,?7)"
)

// los posts solo lo pueden hacer maestros o administradores
func CreatePost(post core.ApiPostPublication, id int) (err error) {
	var database = openDB()
	defer database.Close()

	_, err = database.Exec(createPostQuery,
		post.Title, post.Content, post.Mineature, id, post.Topic, time.Now().Unix(), post.Introduction,
	)

	return
}

// no requiere mucha explicacion , pero tambien se requiere usar en el frontend
func GetPost(id int) (post core.ApiGetPublication) {
	db := openDB()
	defer db.Close()
	db.QueryRow("SELECT * FROM publications WHERE ID=?1", id).Scan(&post.ID, &post.Title, &post.Mineature, &post.Content, &post.Author, &post.Topic, &post.Date, &post.Introduction)
	return
}

// esto es para lo que seria la pagina principal donde se pueden ver otros proyectos y otras cosas
// esto se va a usar en el frontend
func GetPosts(page int, topic string) (posts []core.ApiGetPublication) {
	db := openDB()

	defer db.Close()
	id := PublicationsGetElement(topic, page)
	rows, err := db.Query("SELECT id,title,mineature,author,datePublication,introduction FROM publications WHERE ID<=?1 ORDER BY ID DESC LIMIT ?2", id, core.PostPerPage)
	if err != nil {
		fmt.Println("someting is wrong")
	}
	for rows.Next() {
		var post core.ApiGetPublication
		rows.Scan(&post.ID, &post.Title, &post.Mineature, &post.Author, &post.Date, &post.Introduction)
		posts = append(posts, post)
	}
	return
}

func GetPostsRows(page int, topic string) (rows *sql.Rows) {
	db := openDB()

	defer db.Close()
	id := PublicationsGetElement(topic, page)
	rows, err := db.Query("SELECT id,title,mineature,author,datePublication,introduction FROM publications WHERE ID<=?1 ORDER BY ID DESC LIMIT ?2", id, core.PostPerPage)
	if err != nil {
		fmt.Println("someting is wrong")
	}

	return
}

// es para la api , al momento de buscar algun tema
func GetTopics() (topics []string) {
	db := openDB()
	defer db.Close()
	rows, _ := db.Query("SELECT DISTINCT topic FROM publications")
	for rows.Next() {
		topic := ""
		rows.Scan(&topic)
		topics = append(topics, topic)
	}
	return
}

// obtengo la cantidad de publicaciones, solo sirve para poder mostrar la pagina
func PublicationsSize(topic string) (size int) {
	db := openDB()
	defer db.Close()

	db.QueryRow("SELECT COUNT * FROM publications WHERE topic=?1 OR  \"any\"=?1", topic).Scan(&size)

	return
}
func PublicationsGetElement(topic string, page int) (idPage int) {
	db := openDB()
	defer db.Close()
	// creo que deberia de funcionar este query

	db.QueryRow("SELECT id FROM publications  WHERE topic=?1 OR \"any\"=?1 ORDER BY ID DESC", topic).Scan(&idPage)
	return idPage
}
