package _006_Zig_Zag_Conversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	result := convert("PAYPALISHIRING", 3)
	res := "PAHNAPLSIIGYIR"
	if res != result {
		t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
	}

	result = convert("PAYPALISHIRING", 4)
	res = "PINALSIGYAHRPI"
	if res != result {
		t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
	}
}
