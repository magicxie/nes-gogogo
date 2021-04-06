package bus

import (
	. "nes6502/io"
	"nes6502/misc"
	. "nes6502/ppu"
	. "nes6502/ram"
	. "nes6502/rom"
)

type RamMapper struct {
	Ram         *Ram
	PpuRegister *PPURegister
	IoRegister  *IORegisters
	Rom         *Rom
	SRam        *SRam
	Dma         *DMA
}

func addressInRange(address uint16, lo uint16, hi uint16) bool {
	return (address >= lo) && (address < hi)
}

/**
+---------+-------+-------+-----------------------+
    | 地址    | 大小  | 标记  |         描述          |
    +---------+-------+-------+-----------------------+
    | $0000   | $800  |       | RAM                   |
    | $0800   | $800  | M     | RAM                   |
    | $1000   | $800  | M     | RAM                   |
    | $1800   | $800  | M     | RAM                   |
    | $2000   | 8     |       | Registers             |
    | $2008   | $1FF8 |  R    | Registers             |
    | $4000   | $20   |       | Registers             |
    | $4020   | $1FDF |       | Expansion ROM         |
    | $6000   | $2000 |       | SRAM                  |
    | $8000   | $4000 |       | PRG-ROM               |
    | $C000   | $4000 |       | PRG-ROM               |
    +---------+-------+-------+-----------------------+
*/
func (ramMapper *RamMapper) ReadByte(address uint16) byte {
	if address < 0x2000 {
		return ramMapper.Ram.ReadByte(address % 0x0800)
	}
	//$2000-$2007	$0008	NES PPU registers
	if addressInRange(address, 0x2000, 0x4000) {
		misc.Console.Trace("\nPPU reg Access %04X", address%8)
		if address%8+uint16(1) > uint16(ramMapper.PpuRegister.Memory.Size()) {
			misc.Console.Error("\nppu reg overflow. size:%X,addr:%X\n", ramMapper.PpuRegister.Memory.Size(), address)
		}

		return ramMapper.PpuRegister.ReadByte(address % 8)
	}
	//IO register
	if addressInRange(address, 0x4000, 0x4020) {
		return ramMapper.IoRegister.ReadByte(address - 0x4000)
	}
	//Expansion ROM
	if addressInRange(address, 0x4020, 0x6000) {

	}
	if addressInRange(address, 0x6000, 0x8000) {

	}
	if addressInRange(address, 0x8000, 0xFFFF) {
		return ramMapper.Rom.Program.ReadByte(address % (0x4000 * uint16(ramMapper.Rom.Header.PGMMirrors)))

	}
	return 0
}

func (ramMapper *RamMapper) WriteByte(address uint16, data byte) {
	if address < 0x2000 {
		ramMapper.Ram.WriteByte(address%0x0800, data)
		ramMapper.Ram.WriteByte(address%0x0800+0x0800, data)
		ramMapper.Ram.WriteByte(address%0x0800+0x1000, data)
		ramMapper.Ram.WriteByte(address%0x0800+0x1800, data)
	}
	if addressInRange(address, 0x2000, 0x4000) {
		ramMapper.PpuRegister.WriteByte(address%8, data)

		/**
		$2000-$2007	$0008	NES PPU registers
		$2008-$3FFF	$1FF8	Mirrors of $2000-2007 (repeats every 8 bytes)
		$4000-$4017	$0018	NES APU and I/O registers
		*/

	}

	//IO register
	if addressInRange(address, 0x4000, 0x4020) {

		ramMapper.IoRegister.WriteByte(address-0x4000, data)

		//TODO move to cycles
		if address == DMA_ADDRESS {
			ramMapper.Dma.Copy(data)
		}
	}

	//Expansion ROM         |
	if addressInRange(address, 0x4020, 0x6000) {

	}

	//SRAM
	if addressInRange(address, 0x6000, 0x8000) {

	}

	//PRG-ROM
	if addressInRange(address, 0x8000, 0xC000) {

	}
	if addressInRange(address, 0xC000, 0xF000) {

	}
}
