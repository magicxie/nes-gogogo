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
	if len(data) == 1 {
		ram.data[address] = data[0]
	}
	if len(data) == 2 {
		ram.data[address] = data[1]
		ram.data[address+1] = data[0]
	}
}

func (ram *Ram) ReadBytes(address uint16, bytes int) []byte {
	if bytes == 1 {
		return []byte{ram.data[address]}
	} else {
		//$LLHH
		return []byte{ram.data[address+1], ram.data[address]}
	}
}

func (ram *Ram) WriteUint16(address uint16, data uint16) {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[0:], data)
	for k, v := range b {
		ram.data[address+uint16(k)] = v
	}
}

func (ram *Ram) ReadByte(address uint16) byte {
	return ram.data[address]
}

//USELESS?
func (ram *Ram) ReadUint16(address uint16) uint16 {
	return binary.LittleEndian.Uint16(ram.data[address : address+2])
}
