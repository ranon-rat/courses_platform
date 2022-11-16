package db

import (
	"fmt"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
)

/*
-- sqlite
-- el admin puede agregar a profesores y a nuevos admins
-- pero debe de agregarse manualmente
-- con esto tambien se puede guardar los comentarios
CREATE TABLE publications(
    ID INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    mineature TEXT NOT NULL,--url de un servicio de hosting
    content TEXT NOT NULL,--con esto solo guardare el contenido en texto
    author INTEGER NOT NULL,-- solo usare el id del usuario que publico esto para poder cargar algunas cosas,
    -- en si solo es importante ver lo que son las publicaciones asi que con esto solo se debe de checar
    -- las publicaciones que tengan el id de un admin
    topic VARCHAR(64),--no croe usar mar de esta cantidad de caracteres , de hecho me parece demasiado pero puede que tenga la necesidad
    datePublication integer--tal vez lo vaya a cambiar
);

*/
//los posts solo lo pueden hacer maestros o administradores
func NewPost(post core.ApiPostPublication, id int) {
	db := openDB()
	defer db.Close()
	db.Exec(`INSERT INTO publications
			 	(title,mineature,content,topic,author,datePublication) 
			VALUES(?1,?2,?3,?4,?5,?6)`,
		post.Title, post.Mineature, post.Content, post.Topic, id, time.Now().Unix())
}

// no requiere mucha explicacion , pero tambien se requiere usar en el frontend
func GetPost(id int) (post core.ApiGetPublication) {
	db := openDB()
	defer db.Close()
	db.QueryRow("SELECT * FROM publications WHERE ID=?1", id).Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.Topic, &post.Date)
	return
}

// esto es para lo que seria la pagina principal donde se pueden ver otros proyectos y otras cosas
// esto se va a usar en el frontend
func GetPosts(page int, topic string) (posts []core.ApiGetPublication) {
	db := openDB()

	defer db.Close()
	size := PublicationsSize(topic)
	rows, err := db.Query("SELECT (ID,title,mineature,author,datePublication) FROM publications WHERE ID<=?1 AND ID>=?2 ORDER BY DESC", (size - (page * core.PostPerPage)), size-(page*core.PostPerPage)+core.PostPerPage+1)
	if err != nil {
		fmt.Println("someting is wrong")

	}
	for rows.Next() {
		var post core.ApiGetPublication
		rows.Scan(&post.ID, &post.Title, &post.Author, &post.Date)
		posts = append(posts, post)
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

// obtengo la cantidad de publicaciones, me sirve para poder hacer algunas cosas en general
func PublicationsSize(topic string) (size int) {
	db := openDB()
	defer db.Close()

	if topic == "any" {
		db.QueryRow("SELECT COUNT(*) FROM publications").Scan(&size)
		return
	}
	db.QueryRow("SELECT COUNT(*) FROM publications WHERE topic=?1", topic).Scan(&size)

	return
}
