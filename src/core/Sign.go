package core

/*
CREATE TABLE users(
    ID iNTEGER PRIMARY KEY,
    privileges INT NOT NULL,--1:admin,2:teacher,3:pupil
    username VARCHAR(64) NOT NULL UNIQUE,--es innecesario tener mas
    pass VARCHAR(64) NOT NULL,--sha256(password+token)
    token INTEGER NOT NULL , --tiempo de registro+numero aleatorio del 1 al 1000
    email VARCHAR(64) NOT NULL UNIQUE, --deberia de encriptar esto pero posiblemente en caso de que se les
    ssid VARCHAR(64) --sha256(password+username+email+id+token+unix-time)
);
*/

// voy a usar el ssid para poder verificar esto
type SignUp struct {
	Privileges int    `json:"privileges"` //por default es el 3
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}

// log in
type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
