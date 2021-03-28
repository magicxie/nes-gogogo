package addressing

import (
	"encoding/binary"
	"fmt"
	. "nes6502/bus"
	. "nes6502/cpu/register"
)

type AddressMode struct {
	Resolve func(opc []byte, bus Bus, register Register) (d byte, addr uint16)
}

var (
	//#	immediate	OPC #$BB	operand is byte BB
	Immediate = AddressMode{Resolve: func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		fmt.Printf("#$%X", opc[0])
		return opc[0], register.PC + 1
	}}
	//zpg	zeropage	OPC $LL	operand is zeropage address (hi-byte is zero, address = $00LL)
	Zeropage = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		address := uint16(opc[0])
		b := <-bus.ReadByte(address)
		fmt.Printf("$00%X => %X", address, b)
		return b[0], address
	}}
	//zpg,X	zeropage, X-indexed	OPC $LL,X	operand is zeropage address; effective address is address incremented by X without carry **
	ZeropageX = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		address := uint16(opc[0]+register.X) & 0x00FF
		b := <-bus.ReadByte(address)
		fmt.Printf("%X,%X => $%X = %X\n", opc[0], register.X, address, b)
		return b[0], address
	}}
	//zpg,Y	zeropage, Y-indexed	OPC $LL,Y	operand is zeropage address; effective address is address incremented by Y without carry **
	ZeropageY = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		address := uint16(opc[0]+register.Y) & 0x00FF
		b := <-bus.ReadByte(address)
		fmt.Printf("%X,%X => $%X = %X\n", opc[0], register.Y, address, b)
		return b[0], address
	}}
	//X,ind	X-indexed, indirect	OPC ($LL,X)	operand is zeropage address; effective address is word in (LL + X, LL + X + 1), inc. without carry: C.w($00LL + X)
	IndirectX = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		indAddress := uint16(opc[0]+register.X) & 0x00FF
		address := <-bus.ReadWord(indAddress)
		intAddress := binary.BigEndian.Uint16(address)
		data := <-bus.ReadByte(intAddress)

		fmt.Printf("($%X,%X) => $%X = %X\n", opc[0], register.X, intAddress, data[0])

		return data[0], intAddress
	}}
	//ind,Y	indirect, Y-indexed	OPC ($LL),Y	operand is zeropage address; effective address is word in (LL, LL + 1) incremented by Y with carry: C.w($00LL) + Y
	IndirectY = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		indAddress := uint16(opc[0]+register.Y) & 0x00FF
		address := <-bus.ReadWord(indAddress)
		intAddress := binary.BigEndian.Uint16(address)
		data := <-bus.ReadByte(intAddress)

		fmt.Printf("($%X,%X) => $%X = %X\n", opc[0], register.Y, intAddress, data[0])

		return data[0], intAddress
	}}
	//rel	relative	OPC $BB	branch target is PC + signed offset BB ***
	Relative = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		fmt.Printf("$%X", int16(opc[0]))
		return opc[0], uint16(int16(register.PC) + int16(opc[0]))
	}}
	//abs	absolute	OPC $LLHH	operand is address $HHLL *
	Absolute = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		addr = binary.BigEndian.Uint16(opc)
		data := <-bus.ReadWord(addr)
		fmt.Printf("$%X => $%X = %X\n", binary.BigEndian.Uint16(opc), addr, data[0])
		return data[0], addr
	}}
	//abs,X	absolute, X-indexed	OPC $LLHH,X	operand is address; effective address is address incremented by X with carry **
	AbsoluteX = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		addr = (binary.BigEndian.Uint16(opc) + uint16(register.X)) & 0xFFFF
		data := <-bus.ReadWord(addr)
		fmt.Printf("$%X,%X => $%X = %X\n", binary.BigEndian.Uint16(opc), uint16(register.X), addr, data[0])
		return data[0], addr
	}}
	//abs,Y	absolute, Y-indexed	OPC $LLHH,Y	operand is address; effective address is address incremented by Y with carry **
	AbsoluteY = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		addr = (binary.BigEndian.Uint16(opc) + uint16(register.Y)) & 0xFFFF
		data := <-bus.ReadWord(addr)
		fmt.Printf("$%X,%X => $%X = %X\n", binary.BigEndian.Uint16(opc), uint16(register.Y), addr, data[0])
		return data[0], addr
	}}
	//A	Accumulator	OPC A	operand is AC (implied single byte instruction)
	Accumulator = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		return register.AC, 0
	}}
	//impl	implied	OPC	operand implied
	Implied = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		//fmt.Println("Implied" )
		return 0, register.PC
	}}
	//ind	indirect	OPC ($LLHH)	operand is address; effective address is contents of word at address: C.w($HHLL)
	Indirect = AddressMode{func(opc []byte, bus Bus, register Register) (d byte, addr uint16) {
		indAddress := binary.BigEndian.Uint16(opc)
		address := <-bus.ReadWord(indAddress)
		intAddress := binary.BigEndian.Uint16(address)
		data := <-bus.ReadByte(intAddress)
		fmt.Printf("($%X) => $%X = %X\n", indAddress, intAddress, data[0])
		return data[0], intAddress
	}}
)
