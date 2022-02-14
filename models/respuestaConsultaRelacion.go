package models

/*RespuestaConsultaRelacion tiene el true o false que se obtiene al consultar la relacion entre 2 usuarios*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
