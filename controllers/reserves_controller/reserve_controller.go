package reserves_controller

import (
	"ReserveForm/commons/injections"
	"ReserveForm/controllers/users_controller"
	"ReserveForm/models/reserves"
	"ReserveForm/repositories"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type DataBase struct {
	Type injections.Type
}

func (db DataBase) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user, err := users_controller.User(r, db.Type)
		if err != nil {
			http.Redirect(w, r, "/users/new", http.StatusFound)
			return
		}
		reserves := repositories.Reserve(db.Type).Index(*user)
		users_controller.Show(w, r, db.Type, "reserves/index", reserves)
	}
}

func (db DataBase) New(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, err := users_controller.User(r, db.Type)
		if err != nil {
			http.Redirect(w, r, "/users/new", http.StatusFound)
			return
		}
		places := repositories.Content(db.Type).Index()
		users_controller.Show(w, r, db.Type, "reserves/new", places)

	case "POST":
		date := r.PostFormValue("date")
		content, er := strconv.ParseUint(r.FormValue("ContentID"), 10, 64)
		if er != nil {
			http.Redirect(w, r, "/reserves/new", http.StatusFound)
			return
		}

		user, err := users_controller.User(r, db.Type)
		if err != nil {
			http.Redirect(w, r, "/users/new", http.StatusFound)
			return
		}
		reserve := reserves.Reserve{
			UserID:    user.ID,
			ContentID: uint(content),
			Date:      date,
		}
		repositories.Reserve(db.Type).Create(reserve)
		http.Redirect(w, r, "/reserves", http.StatusFound)
	}
}

func (db DataBase) Delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user, err := users_controller.User(r, db.Type)
		if err != nil {
			http.Redirect(w, r, "/users/new", http.StatusFound)
			return
		}

		vars := mux.Vars(r)
		id, er := strconv.Atoi(vars["id"])
		if er != nil {
			http.Redirect(w, r, "/reserves", http.StatusFound)
			return
		}

		reserve := reserves.Reserve{
			Model:  gorm.Model{ID: uint(id)},
			UserID: user.ID}
		repositories.Reserve(db.Type).Delete(reserve)
		http.Redirect(w, r, "/reserves", http.StatusFound)
	}
}
