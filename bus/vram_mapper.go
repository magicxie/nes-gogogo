package bus

import (
	. "nes6502/ppu/vram"
)

type VRamMapper struct {
	PatternTables *PatternTables
	NameTables    *NameTables
	ImagePalette  *Palette
	SpritePalette *Palette
}

func (vramMapper *VRamMapper) Read(address uint16, bytes int) []byte {
	if address < 0x2000 {
		return vramMapper.PatternTables.ReadBytes(address, bytes)
	}

	if address >= 0x2000 && address < 0x3F00 {
		return vramMapper.NameTables.ReadBytes(address-0x2000, bytes)
	}

	if address >= 0x3F00 && address < 0x3F10 {
		return vramMapper.ImagePalette.ReadBytes(address-0x3F00, bytes)
	}

	if address >= 0x3F10 && address < 0x3F20 {
		return vramMapper.SpritePalette.ReadBytes(address-0x3F10, bytes)
	}

	//Mirrors
	if address >= 0x3F20 && address < 0x4000 {
		return vramMapper.Read((address%0x0020)+0x3F20, bytes)
	}

	if address >= 0x4000 && address <= 0xFFFF {
		return vramMapper.Read(address-0x4000, bytes)
	}

	return []byte{}
}

func (vramMapper *VRamMapper) Write(address uint16, data []byte) {
	if address < 0x2000 {
		vramMapper.PatternTables.WriteBytes(address, data)
	}
	if address >= 0x2000 && address < 0x3F00 {
		vramMapper.NameTables.WriteBytes(address-0x2000, data)
	}

	if address >= 0x3F00 && address < 0x3F10 {
		vramMapper.ImagePalette.WriteBytes(address-0x3F00, data)
	}

	if address >= 0x3F10 && address < 0x3F20 {
		vramMapper.SpritePalette.WriteBytes(address-0x3F10, data)
	}

	if address >= 0x3F20 && address <= 0xFFFF {
		//NOTHING to DO with mirrors
	}

}
