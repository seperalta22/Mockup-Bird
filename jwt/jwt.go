package jwt

import (
	"time"

	"github.com/SergioPeralta22/Mockup-Bird/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//*GeneroJWT genera el encriptado con JWT

func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("EzioPaladini")

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellido":        t.Apellido,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioWeb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
