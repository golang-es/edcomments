package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

// SetRealtimeRouter ruta para el realtime
func SetRealtimeRouter(router *mux.Router) {
	mel := melody.New()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})

	mel.HandleMessage(func(s *melody.Session, msg []byte) {
		mel.Broadcast(msg)
	})
}
