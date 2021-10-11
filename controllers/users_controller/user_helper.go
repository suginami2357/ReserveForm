package users_controller

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/templates"
	"ReserveForm/models/users"
	"ReserveForm/repositories"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

const TOKEN string = "token"
const ID string = "user_id"
const SESSION_NAME string = "ReserveForm_user"

var store *sessions.CookieStore

func Login(w http.ResponseWriter, r *http.Request, t injections.Type, u *users.User) {
	login(0, w, r, t, u) //ログイン有効期限：ブラウザが開かれている間
}

func Login_RememberMe(w http.ResponseWriter, r *http.Request, t injections.Type, u *users.User) {
	login(1209600, w, r, t, u) //ログイン有効期限：２週間
}

func Logout(w http.ResponseWriter, r *http.Request, t injections.Type) {
	ses, _ := store.Get(r, SESSION_NAME)
	id, err := convUint(ses.Values[ID])
	if err != nil {
		return
	}

	ses.Values[ID] = nil
	ses.Values[TOKEN] = nil
	ses.Save(r, w)

	user, _ := repositories.User(t).Take(*users.New_id(id))
	user.Token = nil
	repositories.User(t).Update(*user)
}

func User(r *http.Request, t injections.Type) (*users.User, error) {
	if store == nil {
		return nil, errors.New("user is logout")
	}

	ses, _ := store.Get(r, SESSION_NAME)
	id, err := convUint(ses.Values[ID])
	if err != nil {
		return nil, err
	}

	user, err := repositories.User(t).Take(*users.New_id(id))
	if err != nil {
		return nil, err
	}

	token := ses.Values[TOKEN]
	err = bcrypt.CompareHashAndPassword(user.Token, []byte(token.(string)))
	return user, err
}

func Show(w http.ResponseWriter, r *http.Request, t injections.Type, path string, data interface{}) {
	if user, _ := User(r, t); user != nil {
		templates.Show_Login(w, t, path, data)
	} else {
		templates.Show_Logout(w, t, path, data)
	}
}

//private
func login(age int, w http.ResponseWriter, r *http.Request, t injections.Type, u *users.User) {
	store = sessions.NewCookieStore(random(16))
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   age,
		Secure:   true, //trueで HTTPS通信のみサーバーへ送信する
		HttpOnly: true, //trueで JavaScript からの操作を禁止する
	}

	//認証トークンを生成
	token := base64.URLEncoding.EncodeToString(random(16))

	ses, _ := store.Get(r, SESSION_NAME)
	ses.Values[ID] = u.ID
	ses.Values[TOKEN] = token
	ses.Save(r, w)

	user, _ := repositories.User(t).Take(*u)
	user.Token, _ = bcrypt.GenerateFromPassword([]byte(token), 4)
	repositories.User(t).Update(*user)
}

func random(size uint) []byte {
	value := make([]byte, size)
	rand.Read(value)
	return value
}

func convUint(src interface{}) (uint, error) {
	switch i := src.(type) {
	case uint:
		return i, nil
	default:
		return 0, errors.New("id is not integer")
	}
}
