package io

type Memory struct {
	data []byte
}

func (memory *Memory) Fill(pos uint16, data []byte) {
	for _, d := range data {
		memory.data[pos] = d
	}
}

func (memory *Memory) Size() int {
	return len(memory.data)
}

func (memory *Memory) Allocate(size int) {
	memory.data = make([]byte, size)
}

func (memory *Memory) ReadBytes(address uint16, bytes int) []byte {
	return memory.data[address : address+uint16(bytes)]
}

func (memory *Memory) WriteBytes(address uint16, data []byte) {
	for i, d := range data {
		memory.WriteByte(address+uint16(i), d)
	}
}

func (memory *Memory) ReadByte(address uint16) byte {
	return memory.ReadBytes(address, 1)[0]
}

func (memory *Memory) WriteByte(address uint16, data byte) {
	memory.data[address] = data
}
