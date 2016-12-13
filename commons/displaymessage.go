package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-es/edcomments/models"
)

// DisplayMessage devuelve un mensaje al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
