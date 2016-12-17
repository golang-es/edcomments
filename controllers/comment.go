package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/golang-es/edcomments/models"
	"github.com/golang-es/edcomments/configuration"
	"github.com/golang-es/edcomments/commons"
	"strconv"
	"log"
)

// CommentCreate permite registrar un comentario
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	user := models.User{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}
	comment.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado con éxito"
	commons.DisplayMessage(w, m)
}

// CommentGetAll obtiene todos los comentarios
func CommentGetAll(w http.ResponseWriter, r *http.Request)  {
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	vote := models.Vote{}

	user, _ = r.Context().Value("user").(models.User)
	vars := r.URL.Query()

	db := configuration.GetConnection()
	defer db.Close()

	cComment := db.Where("parent_id = 0")
	if order, ok := vars["order"]; ok {
		if order[0] == "votes" {
			cComment = cComment.Order("votes desc, created_at desc")
		}
	} else {
		if idlimit, ok := vars["idlimit"]; ok {
			registerByPage := 30
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("Error:", err)
			}
			cComment = cComment.Where("id BETWEEN ? AND ?", offset - registerByPage, offset)
		}
		cComment = cComment.Order("id desc")
	}

	cComment.Find(&comments)

	for i := range comments {
		db.Model(&comments[i]).Related(&comments[i].User)
		comments[i].User[0].Password = ""
		comments[i].Children = commentGetChildren(comments[i].ID)

		// Se busca el voto del usuario en sesión
		vote.CommentID = comments[i].ID
		vote.UserID = user.ID
		count := db.Where(&vote).Find(&vote).RowsAffected
		if count > 0 {
			if vote.Value {
				comments[i].HasVote = 1
			} else {
				comments[i].HasVote = -1
			}
		}
	}

	j, err := json.Marshal(comments)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir los comentarios en json"
		commons.DisplayMessage(w, m)
		return
	}

	if len(comments) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.Code = http.StatusNoContent
		m.Message = "No se encontraron comentrarios"
		commons.DisplayMessage(w, m)
	}
}

func commentGetChildren(id uint) (children []models.Comment)  {
	db := configuration.GetConnection()
	defer db.Close()

	db.Where("parent_id = ?", id).Find(&children)
	for i := range children {
		db.Model(&children[i]).Related(&children[i].User)
		children[i].User[0].Password = ""
	}
	return
}
