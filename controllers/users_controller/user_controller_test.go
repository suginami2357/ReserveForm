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

func TestNew(t *testing.T) {
	user1, _ := users.New("sample1@gmail.com", "password1")
	user2, _ := users.New("sample2@gmail.com", "password2")
	user3, _ := users.New("sample3@gmail.com", "password3")
	users_repository.Data = append(users_repository.Data, *user1)
	users_repository.Data = append(users_repository.Data, *user2)
	users_repository.Data = append(users_repository.Data, *user3)

	go servers.Start(injections.Test)
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		t.Error(err)
		return
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		t.Error(err)
		return
	}

	page.Navigate("http://localhost:8000/users/login")

	email := page.FindByID("email")
	email.Fill("sample1@gmail.com")
	time.Sleep(1 * time.Second)

	password := page.FindByID("password")
	password.Fill("password1")
	time.Sleep(1 * time.Second)
	create := page.FindByID("create")
	create.Click()

	time.Sleep(3 * time.Second)
}
