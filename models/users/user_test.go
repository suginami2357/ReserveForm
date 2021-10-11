package users

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("test@yahoo.ne.jp", "password")
	if err != nil {
		t.Fatal(err)
	}
}

//---メールアドレス---
//[Local]@[3LD].[2LD].[TLD]

//空白
func TestNew_Email_Empty(t *testing.T) {
	_, err := New("", "password")
	if err == nil {
		t.Fatal("Email_Empty")
	}
}

//最小文字数
func TestNew_Email_Minimum(t *testing.T) {
	_, err := New("l@d", "password")
	if err == nil {
		t.Fatal("Email_Minimum")
	}
}

//最大文字数
func TestNew_Email_Maximum(t *testing.T) {
	v := string(bytes.Repeat([]byte("a"), 245))
	v = v + "@yahoo.ne.jp"
	_, err := New(v, "password")
	if err == nil {
		t.Fatal("Email_Maximum")
	}
}

//[Local]禁則文字
func TestNew_Email_Local_ProhibitedCharacters(t *testing.T) {
	_, err := New("te!st@yahoo.ne.jp", "password")
	if err == nil {
		t.Fatal("Email_Local_ProhibitedCharacters")
	}
}

//[3LD]に禁則文字
func TestNew_Email_3LD_ProhibitedCharacters(t *testing.T) {
	_, err := New("test@ya!hoo.ne.jp", "password")
	if err == nil {
		t.Fatal("Email_3LD_ProhibitedCharacters")
	}
}

//[2LD]に禁則文字
func TestNew_Email_2LD_ProhibitedCharacters(t *testing.T) {
	_, err := New("test@yahoo.n!e.jp", "password")
	if err == nil {
		t.Fatal("Email_2LD_ProhibitedCharacters")
	}
}

//[TLD]に禁則文字
func TestNew_Email_1LD_ProhibitedCharacters(t *testing.T) {
	_, err := New("test@yahoo.ne.j!p", "password")
	if err == nil {
		t.Fatal("Email_1LD_ProhibitedCharacters")
	}
}

//[Local]@@[3LD].[2LD].[TLD]
func TestNew_Email_At_Excess(t *testing.T) {
	_, err := New("test@@yahoo.ne.jp", "password")
	if err == nil {
		t.Fatal("Email_At_Excess")
	}
}

//[Local][3LD].[2LD].@[TLD]
func TestNew_Email_At_Slide(t *testing.T) {
	_, err := New("testyahoo.ne.@jp", "password")
	if err == nil {
		t.Fatal("Email_At_Slide")
	}
}

//[Local][3LD].[2LD].[TLD]
func TestNew_Email_At_NotExist(t *testing.T) {
	_, err := New("testyahoo.ne.jp", "password")
	if err == nil {
		t.Fatal("Email_At_NotExist")
	}
}

//[Local]@
func TestNew_Email_Domain_NotExist(t *testing.T) {
	_, err := New("test@", "password")
	if err == nil {
		t.Fatal("Email_Domain_NotExist")
	}
}

//@[3LD].[2LD].[TLD]
func TestNew_Email_Local_NotExist(t *testing.T) {
	_, err := New("@yahoo.ne.jp", "password")
	if err == nil {
		t.Fatal("Email_Local_NotExist")
	}
}

//[Local]@[3LD][2LD][TLD]
func TestNew_Email_Dot_NotExist(t *testing.T) {
	_, err := New("test@yahoonejp", "password")
	if err == nil {
		t.Fatal("Email_Dot_NotExist")
	}
}

//[Local]@[3LD].[2LD]
func TestNew_Email_TDL_NotExist(t *testing.T) {
	_, err := New("test@yahoo.0ne", "password")
	if err == nil {
		t.Fatal("Email_TDL_NotExist")
	}
}

//[Local]@[TLD]
func TestNew_Email_TDL_Only(t *testing.T) {
	_, err := New("test@jp", "password")
	if err == nil {
		t.Fatal("Email_TDL_Only")
	}
}
