package routes

import (
	"github.com/golang-es/edcomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCommentRouter contiene las rutas para la creación y consulta de comentarios
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")
	subRouter.HandleFunc("/", controllers.CommentGetAll).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
