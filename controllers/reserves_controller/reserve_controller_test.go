package reserves_controller_test

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
	"ReserveForm/models/contents"
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	"ReserveForm/repositories"
	"ReserveForm/repositories/tests/contents_repository"
	"ReserveForm/repositories/tests/reserves_repository"
	"ReserveForm/repositories/tests/users_repository"
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

	login(page)

	page.FindByID("index").Click()
	time.Sleep(1 * time.Second)

	assert_url(t, page, "http://localhost:8000/reserves")
}

//予約：正常に動作する
func Test_New(t *testing.T) {
	go servers.Start(injections.Test)
	append_content("テスト1")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()

	login(page)

	contents := repositories.Content(injections.Test).Index()
	page.FindByID("content").Select(contents[0].Name)
	page.FindByID("button").Click()
	page.ConfirmPopup()
	time.Sleep(1 * time.Second)

	assert_element(t, page, "button", 1)
}

//予約：予約内容が未選択
func Test_New_Undecided(t *testing.T) {
	go servers.Start(injections.Test)

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()

	login(page)

	page.FindByID("button").Click()
	page.ConfirmPopup()
	time.Sleep(1 * time.Second)

	assert_url(t, page, "http://localhost:8000/reserves/new")
}

//予約キャンセル：正常に動作する
func Test_Delete(t *testing.T) {
	go servers.Start(injections.Test)
	date := time.Now().Format("2006/01/02")
	append_reserve(1, 1, date, "テスト1")

	driver := agouti.ChromeDriver()
	driver.Start()
	page, _ := driver.NewPage()
	defer driver.Stop()

	login(page)

	page.FindByID("index").Click()
	time.Sleep(1 * time.Second)

	page.FindByID("1").Click()
	page.ConfirmPopup()
	time.Sleep(1 * time.Second)

	assert_element(t, page, "button", 0)
}

//private

//ログイン処理
func login(page *agouti.Page) {
	append_user("sample@gmail.com", "password")
	page.Navigate("http://localhost:8000/users/login")
	page.FindByID("email").Fill("sample@gmail.com")
	page.FindByID("password").Fill("password")
	page.FindByID("submit").Click()
	time.Sleep(1 * time.Second)
}

//テストデータにユーザーを追加する
func append_user(email string, password string) {
	user, _ := users.New(email, password)
	user.ID = uint(len(users_repository.Data) + 1)
	users_repository.Data = append(users_repository.Data, *user)
}

//テストデータに予約内容を追加する
func append_content(name string) {
	content := contents.Content{Name: name}
	content.ID = uint(len(reserves_repository.Data) + 1)
	contents_repository.Data = append(contents_repository.Data, content)
}

func append_reserve(userId uint, contentID uint, date string, name string) {
	reserve := reserves.Reserve{
		UserID:      userId,
		ContentID:   contentID,
		ContentName: name,
		Date:        date,
	}
	reserve.ID = uint(len(reserves_repository.Data) + 1)
	reserves_repository.Data = append(reserves_repository.Data, reserve)
}

func assert_url(t *testing.T, page *agouti.Page, url string) {
	var value, err = page.URL()
	if err != nil {
		t.Fatal(err.Error())
	}
	if value != url {
		t.Fatal()
	}
}

func assert_element(t *testing.T, page *agouti.Page, name string, expected int) {
	var elements, _ = page.FindByName(name).Elements()
	actual := len(elements)
	if actual != expected {
		t.Fatal()
	}
}
