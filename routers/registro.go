package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SergioPeralta22/Mockup-Bird/bd"
	"github.com/SergioPeralta22/Mockup-Bird/models"
)

//* Registro es la funcion para crear en la base de datos el registro del usuario.
func Registro(w, http.ResponseWriter, r *http.Request){

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //* El body es un objeto de solo lectura que se lee una sola vez y se destruye.
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return 
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return 
	}

	if len(t.Password) <8 {
		http.Error(w, "Debe especificar una contraseña de al menos 8 caracteres ", 400)
		return 
	}

	_,encontrado,_ :=bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado ==true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return 
	}
	
	_,status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return 
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario"+err.Error(), 400)
		return 
	}
}