package users_controller

import (
	"ReserveForm/commons/injections"
	"ReserveForm/models/alerts"
	"ReserveForm/models/users"
	"ReserveForm/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type DataBase struct {
	Type injections.Type
}

func (db DataBase) New(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Show(w, r, db.Type, "users/new", nil)

	case "POST":
		email := r.PostFormValue("email")
		user, _ := repositories.User(db.Type).Take(users.User{Email: email})
		if user.Email != "" {
			alert := alerts.New(alerts.Danger, "既に使用されているメールアドレスです。")
			user := users.User{Email: email, Alert: alert}
			Show(w, r, db.Type, "users/new", user)
			return
		}

		password := r.PostFormValue("password")
		user, err := users.New(email, password)
		if err != nil {
			alert := alerts.New(alerts.Danger, err.Error())
			user := users.User{Email: email, Alert: alert}
			Show(w, r, db.Type, "users/new", user)
			return
		}

		user, err = repositories.User(db.Type).Create(*user)
		if err != nil {
			alert := alerts.New(alerts.Danger, err.Error())
			vm := users.User{Email: email, Alert: alert}
			Show(w, r, db.Type, "users/new", vm)
			return
		}

		if r.PostFormValue("remember_me") == "" {
			Login(w, r, db.Type, user)
		} else {
			Login_RememberMe(w, r, db.Type, user)
		}
		http.Redirect(w, r, "/reserves/new", http.StatusFound)
	}
}

func (db DataBase) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, err := User(r, db.Type)
		if err == nil {
			http.Redirect(w, r, "/reserves", http.StatusFound)
			return
		}
		Show(w, r, db.Type, "users/login", nil)

	case "POST":
		email := r.PostFormValue("email")
		user, err := repositories.User(db.Type).Take(users.User{Email: email})
		if err != nil {
			alert := alerts.New(alerts.Danger, "メールアドレスまたはパスワードが等しくありません。")
			vm := users.User{Email: email, Alert: alert}
			Show(w, r, db.Type, "users/login", vm)
			return
		}

		//ログイン認証
		password := r.PostFormValue("password")
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
			alert := alerts.New(alerts.Danger, "メールアドレスまたはパスワードが等しくありません。")
			vm := users.User{Email: email, Alert: alert}
			Show(w, r, db.Type, "users/login", vm)
			return
		}

		if r.PostFormValue("remember_me") == "" {
			Login(w, r, db.Type, user)
		} else {
			Login_RememberMe(w, r, db.Type, user)
		}
		http.Redirect(w, r, "/reserves/new", http.StatusFound)
	}
}

func (db DataBase) Logout(w http.ResponseWriter, r *http.Request) {
	Logout(w, r, db.Type)
	http.Redirect(w, r, "/users/login", http.StatusFound)
}
