package reserves

import (
	"testing"
)

func TestFormat_MMDD(t *testing.T) {
	reserve := Reserve{Date: "2021-01-02"}
	var mmdd = reserve.Format_MMDD()
	if mmdd != "01月02日" {
		t.Fatal("Format_MMDD")
	}
}
