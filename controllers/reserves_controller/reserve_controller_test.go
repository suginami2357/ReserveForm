package reserves_controller_test

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

//予約一覧：正常に動作する
func Test_Index(t *testing.T) {
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

//予約：正常に動作する
func Test_New(t *testing.T) {

}

//予約：予約内容が未選択
func Test_New_Undecided(t *testing.T) {

}

//予約キャンセル：正常に動作する
func Test_Delete(t *testing.T) {

}
