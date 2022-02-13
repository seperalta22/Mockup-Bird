package routers

import (
	"errors"
	"strings"

	"github.com/SergioPeralta22/Mockup-Bird/bd"
	"github.com/SergioPeralta22/Mockup-Bird/models"
	"github.com/dgrijalva/jwt-go"
)

//* Email valor de email usado en todos los endpoints
var Email string

//*IDUsuario es el id devuelto del modelo, que se usara en todos los Endpoints
var IDUsuario string

//* Proceso token procesa el token para extraer sus valores
func ProcesoTokens(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("EzioPaladini")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
