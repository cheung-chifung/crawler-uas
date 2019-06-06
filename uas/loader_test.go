package uas

import "testing"

func Test_Load(t *testing.T) {
	r := Load()
	if len(r) == 0 {
		t.Error("cannot load uas file")
	}
}
