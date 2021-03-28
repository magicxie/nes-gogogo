package ppu
import (
	"fmt"
	"nes6502/ram"
	"testing"
)

func TestDMA_Copy(t *testing.T) {

	ram := &ram.Ram{}
	oam := &Oam{}
	oam.Allocate(256)

	dma := &DMA{ram, oam}
	for i := 0;i<0x2FFF;i++{
		ram.WriteBytes(uint16(i), []byte{byte(i%0xFF)})
	}

	dma.Copy(0x0F)

	expect := byte((0x0F00+0xCC)%0xFF)
	actual := oam.ReadBytes(0xCC, 1)[0]
	if actual != expect{
		t.Fatalf("DMA error, expect:%X, actual:%X", expect, actual)
	}
	fmt.Printf("%d:%v", oam.Size(),oam)


}

