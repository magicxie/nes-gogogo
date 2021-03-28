package register

import (
	"testing"
)

func TestFlagCarry(t *testing.T) {

	r := Register{}
	r.SetStatus(SR_FLAG_Carry, true)

	if r.SR != 1 {
		t.Fatalf("SR_FLAG_Carry error, expect:%d, actual:%d", 1, r.SR)
	}
	t.Logf("Test Flag Carry succ")
}


func TestFlagZero(t *testing.T) {

	r := Register{}
	r.SetStatus(SR_FLAG_Zero, true)

	if r.SR != 2 {
		t.Fatalf("SR_FLAG_Zero error, expect:%d, actual:%d", 2, r.SR)
	}
	t.Logf("Test Flag Zero succ")
}



func TestFlagInterrupt(t *testing.T) {

	r := Register{}
	r.SetStatus(SR_FLAG_Interrupt, true)

	if r.SR != 4 {
		t.Fatalf("SR_FLAG_Interrupt error, expect:%d, actual:%d", 4, r.SR)
	}
	t.Logf("Test Flag Interrupt succ")
}



func TestFlagDecimal(t *testing.T) {

	r := Register{}
	r.SetStatus(SR_FLAG_Decimal, true)

	if r.SR != 8 {
		t.Fatalf("SR_FLAG_Decimal error, expect:%d, actual:%d", 8, r.SR)
	}
	t.Logf("Test Flag Decimal succ")
}
