package ppu

import (
	"testing"
)

func TestPPU(t *testing.T) {
	ppu := &PPU{}

	ppu.Init()
	ppu.PowerOn()
}
