package db

import (
	"github.com/bruh-boys/courses_platform/src/core"
	"github.com/bruh-boys/courses_platform/src/tools"
)

const (
	isAlreadyRegisteredQuery = "SELECT EXISTS(SELECT 1 FROM users WHERE email=?1)"
	isPasswordCorrectQuery   = "SELECT EXISTS(SELECT 1 FROM users WHERE email=?1 AND pass=?2)"
)

const (
	isValidSessionQuery = "SELECT EXISTS(SELECT 1 FROM users WHERE ssid=?1&& email=?2)"
	getSessionPrivQuery = "SELECT privileges,ID FROM users WHERE ssid=?1"
)

const (
	signUpQuery = "INSERT INTO users(privileges,username,email,pass,token) VALUES(?1,?2,?3,?4,?5)"
	signInQuery = "UPDATE users SET ssid=?1 WHERE pass=?2 AND email=?3"
)

const (
	getTokenAndIdQuery = "SELECT token,id FROM users WHERE email=?1"
)

func IsAlreadyRegistered(email string) (ok bool, err error) {
	var database = openDB()
	defer database.Close()

	err = database.QueryRow(isAlreadyRegisteredQuery, email).Scan(&ok)
	return
}

func IsPasswordCorrect(email, pass string) (ok bool, err error) {
	var database = openDB()
	defer database.Close()

	var token int

	if ok, err = IsAlreadyRegistered(email); err != nil || !ok {
		return false, err
	}

	if token, _, err = GetTokenAndId(email); err != nil {
		return false, err
	}

	err = database.QueryRow(
		isPasswordCorrectQuery, email, tools.HashPassword(pass, token),
	).Scan(&ok)

	return
}

func SignUp(data core.SignUp) (err error) {
	var database = openDB()
	defer database.Close()

	token := tools.GenerateToken()
	pass := tools.HashPassword(data.Password, token)

	_, err = database.Exec(
		signUpQuery, 3, data.Username, data.Email, pass, token,
	)

	return
}

func GetTokenAndId(email string) (token, id int, err error) {
	var database = openDB()
	defer database.Close()

	err = database.QueryRow(getTokenAndIdQuery, email).Scan(&token, &id)
	return
}

func SignIn(data core.SignIn) (ssid string, err error) {
	var token, id int

	if token, id, err = GetTokenAndId(data.Email); err != nil {
		return
	}

	var database = openDB()
	defer database.Close()

	ssid = tools.GenerateSSID(id, token, data.Email, data.Password)

	_, err = database.Exec(
		signInQuery, tools.GenerateHash(ssid), tools.HashPassword(data.Password, token), data.Email,
	)

	return
}

func IsValidSesion(email string, ssid string) (valid bool, err error) {
	var database = openDB()
	defer database.Close()

	err = database.QueryRow(isValidSessionQuery, ssid, email).Scan(&valid)
	return
}

func GetSession(ssid string) (priv, id int, err error) {
	var database = openDB()
	defer database.Close()

	err = database.QueryRow(getSessionPrivQuery, tools.GenerateHash(ssid)).Scan(&priv, &id)
	return
}
