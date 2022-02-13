package models

//* RespuestaLogin tiene el campo que se devuelve con el login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
