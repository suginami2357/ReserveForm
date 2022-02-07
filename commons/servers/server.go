package servers

import (
	"ReserveForm/commons/injections"
	"ReserveForm/controllers/reserves_controller"
	"ReserveForm/controllers/users_controller"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start(t injections.Type) {
	//依存性を注入
	var uc = users_controller.DataBase{Type: t}
	var rc = reserves_controller.DataBase{Type: t}

	//RESTful APIのURL設計
	router := mux.NewRouter().StrictSlash(true)

	var root string
	var wd, _ = os.Getwd()
	switch t {
	case injections.Postgres:
		root = wd
	case injections.Test:
		root = wd + "/../../"
	default:
		panic("argument is undefined")
	}

	//js,cssをこのフォルダ配下に配置
	name := "/resources/"
	router.NotFoundHandler = http.StripPrefix(name, http.FileServer(http.Dir(root+name)))

	router.HandleFunc("/users/new", uc.New)
	router.HandleFunc("/users/login", uc.Login)
	router.HandleFunc("/users/logout", uc.Logout)

	router.HandleFunc("/reserves", rc.Index)
	router.HandleFunc("/reserves/new", rc.New)
	router.HandleFunc("/reserves/{id}/delete", rc.Delete)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	// log.Fatal(http.ListenAndServe(":8000", router))
}
