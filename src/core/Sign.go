package core

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
