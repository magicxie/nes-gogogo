package bus

import (
	. "nes6502/io"
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
func (ramMapper *RamMapper) Read(address uint16, bytes int) []byte {
	if address < 0x2000 {
		return ramMapper.Ram.ReadBytes(address%0x0800, bytes)
	}
	//$2000-$2007	$0008	NES PPU registers
	if address >= 0x2000 && address < 0x4000 {
		return ramMapper.PpuRegister.ReadBytes(address%8, bytes)
	}
	//IO register
	if address >= 0x4000 && address < 0x4020 {
		return ramMapper.IoRegister.ReadBytes(address-0x4000, bytes)
	}
	//Expansion ROM
	if address >= 0x4020 && address < 0x6000 {

	}
	if address >= 0x6000 && address < 0x8000 {

	}
	if address >= 0x8000 && address < 0xFFFF {
		return ramMapper.Rom.Program.ReadBytes(address%(0x4000*uint16(ramMapper.Rom.Header.PGMMirrors)), bytes)

	}
	return []byte{}
}

func (ramMapper *RamMapper) Write(address uint16, data []byte) {
	if address < 0x2000 {
		ramMapper.Ram.WriteBytes(address%0x0800, data)
		ramMapper.Ram.WriteBytes(address%0x0800+0x0800, data)
		ramMapper.Ram.WriteBytes(address%0x0800+0x1000, data)
		ramMapper.Ram.WriteBytes(address%0x0800+0x1800, data)
	}
	if address >= 0x2000 && address < 0x3FFF {
		ramMapper.PpuRegister.WriteBytes(address%8, data[0])

		/**
		$2000-$2007	$0008	NES PPU registers
		$2008-$3FFF	$1FF8	Mirrors of $2000-2007 (repeats every 8 bytes)
		$4000-$4017	$0018	NES APU and I/O registers
		*/

	}

	//IO register
	if address >= 0x4000 && address < 0x4020 {

		ramMapper.IoRegister.WriteBytes(address-0x4000, data[0])

		//TODO move to cycles
		if address == DMA_ADDRESS {
			ramMapper.Dma.Copy(data[0])
		}
	}

	//Expansion ROM         |
	if address >= 0x4020 && address < 0x6000 {

	}

	//SRAM
	if address >= 0x6000 && address < 0x8000 {

	}

	//PRG-ROM
	if address >= 0x8000 && address < 0xC000 {

	}
	if address >= 0xC000 && address < 0xF000 {

	}
}
