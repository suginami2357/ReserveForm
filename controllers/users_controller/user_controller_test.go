package users_controller_test

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
	"ReserveForm/models/users"
	"ReserveForm/repositories/tests/users_repository"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

//新規作成：正常に動作する
func Test_New(t *testing.T) {
	go servers.Start(injections.Test)

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/new")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	var url, _ = page.URL()
	if url != "http://localhost:8000/reserves/new" {
		t.Fatal()
	}
}

//新規作成：emailが既に存在する
func Test_New_Email_Exist(t *testing.T) {
	go servers.Start(injections.Test)
	append_user("sample@gmail.com", "password")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/new")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	alert := page.FindByID("alert")
	var message, _ = alert.Text()
	if message != "既に使用されているメールアドレスです。" {
		t.Fatal()
	}
}

//新規作成：emailがメールアドレスの形式でない
func Test_New_Email_InvalidFormat(t *testing.T) {
	go servers.Start(injections.Test)

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/new")
	page.FindByID("email").Fill("'sample'@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	alert := page.FindByID("alert")
	var message, _ = alert.Text()
	if message != "メールアドレスの形式ではありません。" {
		t.Fatal()
	}
}

//新規作成：passwordに使用不可能な文字が使われている
func Test_New_Password_ForbiddenCharacters(t *testing.T) {
	go servers.Start(injections.Test)

	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()
	page.Navigate("http://localhost:8000/users/new")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("Ｐassword")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	alert := page.FindByID("alert")
	var message, _ = alert.Text()
	if message != "パスワードは半角英数字と記号のみ使用可能です。" {
		t.Fatal()
	}
}

//ログイン：正常に動作する
func Test_Login(t *testing.T) {
	go servers.Start(injections.Test)
	append_user("sample@gmail.com", "password")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/login")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	var url, _ = page.URL()
	if url != "http://localhost:8000/reserves/new" {
		t.Fatal()
	}
}

//ログイン：emailが存在しない
func Test_Login_Email_NotFound(t *testing.T) {
	go servers.Start(injections.Test)
	append_user("sample@gmail.com", "password")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/login")
	page.FindByID("email").Fill("Sample@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	alert := page.FindByID("alert")
	var message, _ = alert.Text()
	if message != "メールアドレスまたはパスワードが違います。" {
		t.Fatal()
	}
}

//ログイン：passwordが一致しない
func Test_Login_Password_Mismatch(t *testing.T) {
	go servers.Start(injections.Test)
	append_user("sample@gmail.com", "password")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/login")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("Password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	alert := page.FindByID("alert")
	var message, _ = alert.Text()
	if message != "メールアドレスまたはパスワードが違います。" {
		t.Fatal()
	}
}

//ログアウト：正常に動作する
func Test_Logout(t *testing.T) {
	go servers.Start(injections.Test)
	append_user("sample@gmail.com", "password")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()
	page.Navigate("http://localhost:8000/users/login")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)

	var url, _ = page.URL()
	if url != "http://localhost:8000/reserves/new" {
		t.Fatal()
	}

	page.FindByID("logout").Click()
	time.Sleep(1 * time.Second)

	url, _ = page.URL()
	if url != "http://localhost:8000/users/login" {
		t.Fatal()
	}
}

//private
func append_user(email string, password string) {
	user, _ := users.New(email, password)
	user.ID = uint(len(users_repository.Data) + 1)
	users_repository.Data = append(users_repository.Data, *user)
}
