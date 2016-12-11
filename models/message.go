package models

// Message mensaje para el cliente de la api
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
