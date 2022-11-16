package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bruh-boys/courses_platform/src/core"
)

/*
CREATE TABLE users(
    ID iNTEGER PRIMARY KEY,
    privileges INT NOT NULL,--1:admin,2:teacher,3:pupil
    username VARCHAR(64) NOT NULL UNIQUE,--es innecesario tener mas
    pass VARCHAR(64) NOT NULL,--sha256(password+token)
    token INTEGER NOT NULL , --tiempo de registro+numero aleatorio // solo se puede usar para el momento de entrar al a cuenta, no es algo que sirva de mucho
    email VARCHAR(64) NOT NULL UNIQUE, --deberia de encriptar esto pero posiblemente en caso de que se les
    ssid VARCHAR(64), --sha256(password+email+id+token+unix-time)
    pfp TEXT-- solo es una url
);
*/
// el error seria en caso de que el correo se repitiera o algo parecido
func SignUp(sUp core.SignUp) (err error) {
	db := openDB()

	defer db.Close()

	token := rand.Int() + int(time.Now().Unix())
	password := hashIt(fmt.Sprintf("%s%d", sUp.Password, token))
	_, err = db.Exec("INSERT INTO users(privileges,username,email,pass,token) VALUES(?1,?2,?3,?4,?5)",
		sUp.Privileges,
		sUp.Username,
		sUp.Email,
		password,
		token)
	return
}

// en caso de que el la contraseña no sea correcta debo de checar eso , aun que al no iniciar sesion seria una forma de ver eso, manejare eso por el cliente
func SignIn(sgIn core.SignIn) (ssid string) {
	db := openDB()
	defer db.Close()
	token, id := 0, 0
	db.QueryRow("SELECT (token,id) FROM users WHERE email=?1", sgIn.Email).Scan(&token, &id)

	ssid = hashIt(fmt.Sprintf("%s%s%d%d%d%d", sgIn.Password, sgIn.Email, id, time.Now().Unix(), token, rand.Int()))

	db.Exec("UPDATE users SET ssid=?1 WHERE pass=?2 AND email=?3", hashIt(ssid), hashIt(fmt.Sprintf("%s%d", sgIn.Password, token)), sgIn.Email)
	return

}

// solo checo si el usuario ya a iniciado sesion, utilizo una cookie que se genera automaticamente al iniciar sesion
func Existence(ssid string) (priv, id int) {
	db := openDB()
	defer db.Close()
	r := db.QueryRow("SELECT (privileges,ID) FROM users WHERE ssid=?1", hashIt(ssid))
	r.Scan(&priv, &id)
	return
}
