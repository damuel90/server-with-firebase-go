package errors

//ServiceError se usa para retorna un mensaje de error en el consumo de un servicio
type ServiceError struct {
	Message string `json:"message"`
}
