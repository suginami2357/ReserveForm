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
	_, err := users_controller.User(r, db.Type)
	if err != nil {
		http.Redirect(w, r, "/users/new", http.StatusFound)
		return
	}

	places := repositories.Place(db.Type).Index()
	users_controller.Show(w, r, db.Type, "reserves/new", places)
}

func (db DataBase) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	date := r.PostFormValue("date")
	place, er := strconv.ParseUint(r.FormValue("PlaceID"), 10, 64)
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
		UserID:  user.ID,
		PlaceID: uint(place),
		Date:    date,
	}
	repositories.Reserve(db.Type).Create(reserve)
	http.Redirect(w, r, "/reserves", http.StatusFound)
}

func (db DataBase) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

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
