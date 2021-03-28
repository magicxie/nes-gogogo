package io

type IORegisters struct {
	Size int
	data []byte
}

func (ioRegisters *IORegisters) ReadBytes(address uint16, bytes int) []byte {
	return ioRegisters.data[address : address+1]
}

func (ioRegisters *IORegisters) WriteBytes(address uint16,  data byte) {
	ioRegisters.data[address] = data
}