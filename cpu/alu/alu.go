package alu

import . "nes6502/cpu/register"

type ALU struct {
	bcdMode  bool
	register *Register
}

func (alu *ALU) Init(register *Register) {
	//For NES (A203 CPU) does not have a BCD mode
	alu.bcdMode = false
	alu.register = register
}

func (alu *ALU) Add(a byte, b byte) byte {

	var carry int8 = 0
	if alu.register.GetStatus(SR_FLAG_Carry) {
		carry = 1
	}
	//to singed
	var r int8 = int8(a) + int8(b) + carry

	alu.NegativeOut(r)
	alu.ZeroOut(r)
	alu.Overflow(int16(a) + int16(b) + int16(carry))
	alu.CarryOut(uint16(a) + uint16(b) + uint16(carry))

	return byte(r)
}

func (alu *ALU) IsMinus(a byte) bool {
	return (a >> 7) == 1
}

func (alu *ALU) TwosComplete(a byte) byte {
	if alu.IsMinus(a) {
		return ^a + 1
	} else {
		return a
	}
}

/**
A = A - M - (1-C)
*/
func (alu *ALU) Sub(a byte, b byte) byte {

	var carry int8 = 0
	if alu.register.GetStatus(SR_FLAG_Carry) {
		carry = 1
	}
	//to singed
	var r int8 = int8(a) - int8(b) - carry

	alu.NegativeOut(r)
	alu.ZeroOut(r)
	alu.Overflow(int16(a) - int16(b) - int16(carry))
	alu.CarryOut(uint16(a) - uint16(b) - uint16(carry))

	return byte(r)

}

func (alu *ALU) Or(a byte, b byte) byte {
	r := a | b
	alu.NegativeOut(int8(r))
	alu.ZeroOut(int8(r))
	return r
}

func (alu *ALU) Xor(a byte, b byte) byte {
	r := a ^ b
	alu.NegativeOut(int8(r))
	alu.ZeroOut(int8(r))
	return r
}

func (alu *ALU) And(a byte, b byte) byte {
	r := a & b
	alu.NegativeOut(int8(r))
	alu.ZeroOut(int8(r))
	return r
}

func (alu *ALU) ShiftLeft(a byte) byte {
	r := a << 1
	alu.NegativeOut(int8(r))
	alu.ZeroOut(int8(r))
	alu.CarryOut(uint16(r << 1))
	return r
}

func (alu *ALU) ShiftRight(a byte) byte {
	r := a >> 1
	alu.NegativeOut(int8(r))
	alu.ZeroOut(int8(r))
	alu.CarryOut(uint16(r >> 1))
	return r
}

func (alu *ALU) Overflow(a int16) {
	//Overflow for signed
	alu.register.SetStatus(SR_FLAG_Overflow, (a <= -128) || (a > 127))
}

/**
When the addition result is 0 to 255, the carry is cleared.
When the addition result is greater than 255, the carry is set.
*/
func (alu *ALU) CarryOut(a uint16) {
	//Carry for unsigned
	alu.register.SetStatus(SR_FLAG_Carry, a > 255)
}

func (alu *ALU) ZeroOut(a int8) {
	alu.register.SetStatus(SR_FLAG_Zero, a == 0)
}

func (alu *ALU) NegativeOut(a int8) {
	alu.register.SetStatus(SR_FLAG_Negative, a < 0)
}
