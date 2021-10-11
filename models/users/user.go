package users

import (
	"ReserveForm/models/alerts"
	"errors"
	"regexp"
	"unicode/utf8"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string
	Password []byte
	Token    []byte
	Alert    alerts.Alert
}

func New_id(id uint) *User {
	user := User{}
	user.ID = id
	return &user
}

func New(email string, password string) (*User, error) {
	if utf8.RuneCountInString(email) < 4 {
		return nil, errors.New("メールアドレスは４文字以上必要です。")
	}

	if utf8.RuneCountInString(email) > 254 {
		return nil, errors.New("メールアドレスは最大２５４文字です。")
	}

	//[Local]@[3LD].[2LD].[TLD]
	//◇Local
	//・先頭：^[a-zA-Z0-9_+-]
	//・それ以外：[a-zA-Z0-9_.+-]*
	//※ ! # $ % & ' * + - / = ? ^ _ ` { | } ~ . を使用出来るがプロバイダー側で制限を掛けている様なので一部のみ有効
	//◇3LD・2LD
	//・先頭：[a-zA-Z0-9]
	//・末尾：[a-zA-Z0-9]*
	//・それ以外：[a-zA-Z0-9-]*
	//◇TLD
	//・半角英 [a-zA-Z]{2,}$
	if !regexp.MustCompile(`^[a-zA-Z0-9_+-][a-zA-Z0-9_.+-]*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.){1,2}[a-zA-Z]{2,}$`).
		Match([]byte(email)) {
		return nil, errors.New("メールアドレスの形式ではありません。")
	}

	if utf8.RuneCountInString(password) < 8 {
		return nil, errors.New("パスワードは８文字以上必要です。")
	}

	//ハッシュ化の最大長は50〜72バイト
	//https://dzone.com/articles/be-aware-that-bcrypt-has-a-maximum-password-length
	if utf8.RuneCountInString(password) > 50 {
		return nil, errors.New("パスワードは最大５０文字です。")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9~!@#$%^&*()_+-={}|:";'<>?,./]+$`).Match([]byte(password)) {
		return nil, errors.New("パスワードは半角英数字と記号のみ使用可能です。")
	}

	//第二引数（暗号強度）は4-31 規定値は10
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return &User{Email: email, Password: hashed}, nil
}
