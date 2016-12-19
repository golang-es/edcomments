package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/golang-es/edcomments/controllers"
)

// SetVoteRouter es la ruta para el registro de un voto
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.VoteRegister).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
