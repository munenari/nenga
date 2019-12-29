package model

import "testing"

func TestConvertAddress2EM(t *testing.T) {
	v := "大阪府市石橋2-8-10-D"
	expected := "大阪府市石橋二―八―一〇―Ｄ"
	res := convertAddress2EM(v)
	if expected != res {
		t.Error(v, "converted,", "actual:", res, "expected:", expected)
	}
}
