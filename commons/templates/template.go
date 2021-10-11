package templates

import (
	"ReserveForm/commons/injections"
	"net/http"
	"os"
	"text/template"
)

func Show_Login(w http.ResponseWriter, t injections.Type, path string, data interface{}) {
	root := root(t)
	template := new_template(w, root, path, data)
	template.ParseFiles(root + "/views/navbar_login.html")
	template.Execute(w, data)
}

func Show_Logout(w http.ResponseWriter, t injections.Type, path string, data interface{}) {
	root := root(t)
	template := new_template(w, root, path, data)
	template.ParseFiles(root + "/views/navbar_logout.html")
	template.Execute(w, data)
}

//private
func root(t injections.Type) string {
	var wd, _ = os.Getwd()
	switch t {
	case injections.Sqlite:
		return wd
	case injections.Test:
		return wd + "/../../"
	default:
		panic("argument is undefined")
	}
}

func new_template(w http.ResponseWriter, root string, path string, data interface{}) *template.Template {
	var t = template.Must(template.ParseFiles(root + "/views/" + path + ".html"))
	t.ParseFiles(root + "/views/header.html")
	t.New("javascript").Parse(`<script src="../resources/javascripts/` + path + `.js"></script>`)
	t.New("stylesheet").Parse(`<link rel="stylesheet" href="../resources/stylesheets/` + path + `.css">`)
	t.ParseFiles(root + "/views/alerts/alert.html")
	t.ParseFiles(root + "/views/footer.html")
	return t
}
