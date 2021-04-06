package ram

import (
	"encoding/binary"
)

type Ram struct {
	data [0xffff]byte
}

func (ram *Ram) Dump(data []byte) {
	for i, b := range data {
		ram.data[i] = b
	}
}

func (ram *Ram) ZeroPage(address uint16) byte {
	return ram.data[0x00FF&address]
}

func (ram *Ram) Stack(address uint16) byte {
	return ram.data[0x01FF+address]
}

func (ram *Ram) WriteByte(address uint16, data byte) {
	ram.data[address] = data
}

func (ram *Ram) WriteBytes(address uint16, data []byte) {
	for idx, d := range data {
		ram.data[address+uint16(idx)] = d
	}
}

func (ram *Ram) ReadBytes(address uint16, bytes int) []byte {
	if bytes == 1 {
		return []byte{ram.data[address]}
	} else {
		//$LLHH
		return []byte{ram.data[address], ram.data[address+1]}
	}
}

func (ram *Ram) WriteUint16(address uint16, data uint16) {
	ram.data[address] = byte(data >> 8)
	ram.data[address+1] = byte(data)
}

func (ram *Ram) ReadByte(address uint16) byte {
	return ram.data[address]
}

//USELESS?
func (ram *Ram) ReadUint16(address uint16) uint16 {
	return binary.BigEndian.Uint16(ram.data[address : address+2])
}
