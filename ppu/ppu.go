package ppu

import (
	. "nes6502/clock"
	. "nes6502/ppu/vram"
	"time"
)

type Frame struct {
	pixels []RGB
}

type PPU struct {
	clock *Clock
	cycle int

	PatternTables []PatternTable
	Oam           *Oam
	ColorPalette  []RGB
	VScan         int
	HScan         int
	Frames        int
	nextFrame     *Frame

	fpsControl *FPSControl
}

const (
	HorizontalScanCycles = 341
	VerticalScanCycles   = 262
	FrameLimit           = 60
)

func (ppu *PPU) MapColor(color byte) RGB {
	return ppu.ColorPalette[color]
}

func (ppu *PPU) DrawPixel() {

	if ppu.HScan == 0 {
		if ppu.VScan == 0 {
			ppu.nextFrame = &Frame{make([]RGB, HorizontalScanCycles*VerticalScanCycles)}
		}
	}
	//ppu.nextFrame.scanLine()

}

type FPSControl struct {
	FrameLimit int
	start      int64
}

func (ctrl *FPSControl) Limit(ppu *PPU) {
	if ppu.Frames%ctrl.FrameLimit == 0 {

		now := time.Now()
		nowNS := now.UnixNano()

		intervalNS := nowNS - ctrl.start
		frames := int64(ctrl.FrameLimit) * 10 * 100000000
		fps := frames / intervalNS

		//Low resolution wait
		if intervalNS < time.Second.Nanoseconds() {
			<-time.After(time.Duration(time.Second.Nanoseconds() - intervalNS))
		}
		println("Actual FPS:", fps, "Actual FPS:", frames/(time.Now().UnixNano()-ctrl.start))
		ctrl.start = time.Now().UnixNano()
	}
	ppu.Frames = (ppu.Frames + 1) % FrameLimit
}

func (ppu *PPU) PowerOn() {

}

func (ppu *PPU) Init() {

	ppu.VScan = 0
	ppu.HScan = 0
	ppu.Frames = 0


	ppu.fpsControl = &FPSControl{
		FrameLimit,
		time.Now().UnixNano(),
	}

}


func (ppu *PPU) AcceptClockPulse(cycles chan int){
	for i := range cycles {

		ppu.cycle += i
		for j := 0; j < i; j++ {

			ppu.DrawPixel()

			if ppu.HScan == 0 {
				if ppu.VScan == 0 {
					ppu.fpsControl.Limit(ppu)
				}
				ppu.VScan = (ppu.VScan + 1) % VerticalScanCycles
				//println()
				//print(ppu.VScan," ")
			}
			ppu.HScan = (ppu.HScan + 1) % HorizontalScanCycles
		}
	}

}
