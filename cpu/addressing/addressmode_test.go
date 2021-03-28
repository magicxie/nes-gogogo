package addressing

import (
	. "nes6502/bus"
	. "nes6502/cpu/register"
	. "nes6502/ram"
	"testing"
)

func TestImmidiate(t *testing.T) {
	_bus, _, _register := BeforEach(t)
	//Immediate.Resolve
	oprand := byte(0xFC)
	ll, _ := Immediate.Resolve([]byte{oprand}, *_bus, *_register)
	if ll != oprand {
		t.Fatalf("Immediate error, expect:%d, actual:%d", oprand, ll)
	}

}

func TestIndirect(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_bus.WriteWord(uint16(0x1234), []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := Indirect.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xAB {
		t.Fatalf("Indirect error, expect:%d, actual:%d", 0xAB, ll)
	}

}


func TestIndirectX(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.X = 0xC8
	_bus.WriteWord(uint16(0x12+0xC8)&0xFF, []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := IndirectX.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xAB {
		t.Fatalf("IndirectX error, expect:%d, actual:%d", 0xAB, ll)
	}

}


func TestIndirectY(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.Y = 0x0A
	_bus.WriteWord(uint16(0x12+0x0A)&0xFF, []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := IndirectY.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xAB {
		t.Fatalf("IndirectX error, expect:%d, actual:%d", 0xAB, ll)
	}

}

func TestRelative(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.PC = 0x0841
	_bus.WriteWord(uint16(0x1234), []byte{0x56, 0x78})
	//_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, add := Relative.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0x12 {
		t.Fatalf("Relative error, expect:%d, actual:%d", 0x12, ll)
	}
	if add != (0x0841 + 0x12) {
		t.Fatalf("Relative Address error, expect:%d, actual:%d", (0x0841 + 0x12), ll)
	}

}


func TestZero(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_bus.WriteWord(uint16(0x1234), []byte{0x56, 0x78})
	_bus.WriteByte(0x0012, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := Zeropage.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xAB {
		t.Fatalf("Zero Page error, expect:%d, actual:%d", 0xAB, ll)
	}

}


func TestZeroX(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.X = 0x7A
	_bus.WriteWord(uint16(0x1234), []byte{0x56, 0x78})
	_bus.WriteByte(0x0012+0x7A, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := ZeropageX.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xAB {
		t.Fatalf("Zero Page error, expect:%d, actual:%d", 0xAB, ll)
	}

}


func TestZeroY(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.Y = 0xFF
	_bus.WriteWord(uint16(0x1234), []byte{0x56, 0x78})
	_bus.WriteByte((0x0012+0xFF) & 0xFF, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := ZeropageY.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xAB {
		t.Fatalf("Zero Page error, expect:%d, actual:%d", 0xAB, ll)
	}

}


func TestAbsolute(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.Y = 0x0A
	_bus.WriteWord(uint16(0x1234), []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := Absolute.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0x56 {
		t.Fatalf("Absolute error, expect:%d, actual:%d", 0x56, ll)
	}

}


func TestAbsoluteX(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.X = 0x0A
	_bus.WriteWord(uint16(0x1234+0x0A), []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := AbsoluteX.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0x56 {
		t.Fatalf("AbsoluteX error, expect:%d, actual:%d", 0x56, ll)
	}

}


func TestAbsoluteY(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.Y = 0xFC
	_bus.WriteWord(uint16(0x1234+0xFC), []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := AbsoluteY.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0x56 {
		t.Fatalf("AbsoluteY error, expect:%d, actual:%d", 0x56, ll)
	}

}



func TestAccumulator(t *testing.T) {

	_bus, _, _register := BeforEach(t)
	_register.AC = 0xED
	_bus.WriteWord(uint16(0x1234+0xFC), []byte{0x56, 0x78})
	_bus.WriteByte(0x5678, []byte{0xAB})

	//Immediate.Resolve
	ll, _ := Accumulator.Resolve([]byte{0x12, 0x34}, *_bus, *_register)
	if ll != 0xED {
		t.Fatalf("Accumulator error, expect:%d, actual:%d", 0xED, ll)
	}

}

func BeforEach(t *testing.T) (*Bus, *Ram, *Register) {
	_ram := &Ram{}
	_bus := &Bus{}
	_bus.Init(_ram)
	_register := &Register{}
	_register.Reset()
	return _bus, _ram, _register
}
