package ppu

import (
	"fmt"
	. "nes6502/ram"
)

/**
Direct Memory Access (DMA) is a technique for more efficient copying of data from CPU memory to sprite memory
The whole of sprite memory can be filled by using a single write to $4014
Starting address in CPU memory is specified by the operand for the write multiplied by $100. The 256 bytes starting at this address are copied directly into sprite memory without additional intervention of the CPU
DMA uses memory bus, preventing CPU from using it during this time. This prevents CPU from accessing any more instructions (cycle stealing)
DMA takes the equivalent of about 512 cycles, or about 4.5 scanlines worth
 */
type DMA struct {
	Ram *Ram
	Oam *Oam
}

func (dma *DMA) Copy(address byte) {
	start := uint16(address) << 8
	fmt.Printf("DMA from RAM $0x%04X to OAM\n", start)
	for i := uint16(0); i < 256; i++ {
		dma.Oam.Fill(i, dma.Ram.ReadBytes(start+i, 1))
	}

}

const DMA_ADDRESS = 0x4014
