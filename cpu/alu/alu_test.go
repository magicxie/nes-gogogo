package alu

import (
	. "nes6502/cpu/register"
	"testing"
)

func assertFlag(result byte, expected byte, register *Register, c bool, v bool, n bool, z bool, t *testing.T) {

	if result != expected {
		t.Fatalf("TestALU error, expect:%d, actual:%d", expected, result)
	}

	print("expected ", expected, " ")
	register.PrintStatus()
	if c != register.GetStatus(SR_FLAG_Carry) {
		t.Fatalf("SR_FLAG_Carry error, expect:%v, actual:%v", c, register.GetStatus(SR_FLAG_Carry))
	}
	if v != register.GetStatus(SR_FLAG_Overflow) {
		t.Fatalf("SR_FLAG_Overflow error, expect:%v, actual:%v", v, register.GetStatus(SR_FLAG_Overflow))
	}
	if n != register.GetStatus(SR_FLAG_Negative) {
		t.Fatalf("SR_FLAG_Negative error, expect:%v, actual:%v", n, register.GetStatus(SR_FLAG_Negative))
	}
	if z != register.GetStatus(SR_FLAG_Zero) {
		t.Fatalf("SR_FLAG_Zero error, expect:%v, actual:%v", z, register.GetStatus(SR_FLAG_Zero))
	}
	register.Reset()

}

/**
Tests for ADC
00 + 00 and C=0 gives 00 and N=0 V=0 Z=1 C=0 (simulate)
79 + 00 and C=1 gives 80 and N=1 V=1 Z=0 C=0 (simulate)
24 + 56 and C=0 gives 80 and N=1 V=1 Z=0 C=0 (simulate)
93 + 82 and C=0 gives 75 and N=0 V=1 Z=0 C=1 (simulate)
89 + 76 and C=0 gives 65 and N=0 V=0 Z=0 C=1 (simulate)
89 + 76 and C=1 gives 66 and N=0 V=0 Z=1 C=1 (simulate)
80 + f0 and C=0 gives d0 and N=0 V=1 Z=0 C=1 (simulate)
80 + fa and C=0 gives e0 and N=1 V=0 Z=0 C=1 (simulate)
2f + 4f and C=0 gives 74 and N=0 V=0 Z=0 C=0 (simulate)
6f + 00 and C=1 gives 76 and N=0 V=0 Z=0 C=0 (simulate)
*/
func TestALU_Add(t *testing.T) {
	//fmt.Printf("%b,%b,%b",257, 0x00FF, 257 & 0x00FF)
	register := &Register{}
	alu := &ALU{}
	alu.Init(register)
	assertFlag(alu.Add(0x80, 0xFF, register), 0x7F, register, true, true, false, false, t)
	assertFlag(alu.Add(0x00, 0x00, register), 0x00, register, false, false, false, true, t)

	//register.SetStatus(SR_FLAG_Carry, true)
	assertFlag(alu.Add(0x7F, 0x01, register), 0x80, register, false, true, true, false, t)
	assertFlag(alu.Add(0x7F, 0x80, register), 0xFF, register, false, true, true, false, t)
	assertFlag(alu.Add(0x01, 0xFF, register), 0x00, register, true, true, false, true, t)

	//2f + 4f and C=0 gives 74 and N=0 V=0 Z=0 C=0 (simulate)
	assertFlag(alu.Add(0x80, 0xFA, register), 0xE0, register, true, false, true, false, t)

}
