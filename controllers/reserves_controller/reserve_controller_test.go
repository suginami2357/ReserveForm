package reserves_controller_test

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	"ReserveForm/repositories/tests/reserves_repository"
	"ReserveForm/repositories/tests/users_repository"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

//予約一覧：正常に動作する
func Test_Index(t *testing.T) {
	go servers.Start(injections.Test)
	append_reserve(1, 1)

	//ログイン
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

	page.FindByID("index").Click()
	time.Sleep(1 * time.Second)

	if page.FindByID("1") != nil {
		t.Fatal()
	}
}

//予約：正常に動作する
func Test_New(t *testing.T) {

}

//予約：予約内容が未選択
func Test_New_Undecided(t *testing.T) {

}

//予約キャンセル：正常に動作する
func Test_Delete(t *testing.T) {

}

//private
func append_reserve(userId uint, placeID uint) {
	reserve := reserves.Reserve{UserID: userId, PlaceID: placeID}
	reserve.ID = uint(len(reserves_repository.Data) + 1)
	reserves_repository.Data = append(reserves_repository.Data, reserve)
}

func append_user(email string, password string) {
	user, _ := users.New(email, password)
	user.ID = uint(len(users_repository.Data) + 1)
	users_repository.Data = append(users_repository.Data, *user)
}
