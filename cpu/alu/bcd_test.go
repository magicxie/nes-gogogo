package alu

import (
	"testing"
)

func TestBCD_EncodeDecimal(t *testing.T) {
	bcd := &BCD{}
	s := bcd.EncodeDecimal(1975)
	if 6517 != s {
		t.Fatalf("EncodeDecimal error, expect:%d, actual:%d", 6517, s)
	}

}

func TestBCD_DecodeBcd(t *testing.T) {
	bcd := &BCD{}

	s := bcd.DecodeBcd(6517)

	if 1975 != s {
		t.Fatalf("DecodeBcd error, expect:%d, actual:%d", 1975, s)
	}

}
