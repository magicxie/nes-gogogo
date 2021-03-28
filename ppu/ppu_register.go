package ppu

import . "nes6502/io"

type PPURegister struct {
	Memory
}

/**
端口地址	读写/位	功能描述	解释
$2000	写	PPU控制寄存器	PPUCTRL
-	D1 D0	确定当前使用的名称表	配合微调滚动使用
-	D2	PPU读写显存增量	0(+1 列模式) 1(+32 行模式)
-	D3	精灵用图样表地址	0($0000) 1($1000)
-	D4	背景用图样表地址	0($0000) 1($1000)
-	D5	精灵尺寸(高度)	0(8x8) 1(8x16)
-	D6	PPU 主/从模式	FC没有用到
-	D7	NMI生成使能标志位	1(在VBlank时触发NMI)
*/

func (register *PPURegister) NameTable() byte {
	return register.ReadBytes(0, 1)[0] >> 6
}

func (register *PPURegister) OAMMode() byte {
	return register.ReadBytes(0, 1)[0] << 2 >> 7
}

func (register *PPURegister) SpritePatternAddress() uint16 {
	if register.ReadBytes(0, 1)[0]<<3>>7 == 0 {
		return 0x0000
	} else {
		return 0x1000
	}
}

func (register *PPURegister) BackgroundPatternAddress() uint16 {
	if register.ReadBytes(0, 1)[0]<<4>>7 == 0 {
		return 0x0000
	} else {
		return 0x1000
	}
}

func (register *PPURegister) SpriteHeight() byte {
	if register.ReadBytes(0, 1)[0]<<5>>7 == 0 {
		return 8
	} else {
		return 16
	}
}

func (register *PPURegister) EnableNMI() bool {
	if register.ReadBytes(0, 1)[0]<<7>>7 == 0 {
		return false
	} else {
		return true
	}
}

/**
$2001	写	PPU掩码寄存器	PPUMASK
	-	D0	显示模式	0(彩色) 1(灰阶)
	-	D1	背景掩码	0(不显示最左边那列, 8像素)的背景
	-	D2	精灵掩码	0(不显示最左边那列, 8像素)的精灵
	-	D3	背景显示使能标志位	1(显示背景)
	-	D4	精灵显示使能标志位	1(显示精灵)
	NTSC	D5 D6 D7	颜色强调使能标志位	5-7分别是强调RGB
	PAL	D5 D6 D7	颜色强调使能标志位	5-7分别是强调GRB
*/
func (register *PPURegister) DisplayMode() string {
	if register.ReadBytes(1, 1)[0]>>7 == 0 {
		return "RGB"
	} else {
		return "GRAY"
	}
}

func (register *PPURegister) BackgroundMask() bool {
	return register.ReadBytes(1, 1)[0]<<1>>7 == 0
}

func (register *PPURegister) SpriteMask() bool {
	return register.ReadBytes(1, 1)[0]<<2>>7 == 0
}

func (register *PPURegister) EnableBackground() bool {
	return register.ReadBytes(1, 1)[0]<<3>>7 == 0
}
func (register *PPURegister) EnableSprite() bool {
	return register.ReadBytes(1, 1)[0]<<4>>7 == 0
}

func (register *PPURegister) EnableRGB(NTSC bool) (r bool, g bool, b bool) {
	if NTSC {
		return register.ReadBytes(1, 1)[0]<<5>>7 == 1,
			register.ReadBytes(1, 1)[0]<<6>>7 == 1,
			register.ReadBytes(1, 1)[0]<<7>>7 == 1
	} else {
		return register.ReadBytes(1, 1)[0]<<6>>7 == 1,
			register.ReadBytes(1, 1)[0]<<5>>7 == 1,
			register.ReadBytes(1, 1)[0]<<7>>7 == 1
	}
}

/*

	$2002	读	PPU状态寄存器	PPUSTATUS
	-	D5	精灵溢出标志位	0(当前扫描线精灵个数小于8)
	-	D6	精灵命中测试标志位	1(#0精灵命中) VBlank之后置0
	-	D7	VBlank标志位	VBlank开始时置1, 结束或者读取该字节($2002)后置0
	$2003	写	精灵RAM指针	设置精灵RAM的8位指针
	$2004	读写	精灵RAM数据	读写精灵RAM数据, 访问后指针+1
	$2005	写x2	屏幕滚动偏移	第一个写的值: 垂直滚动 第2个写的值: 水平滚动
	$2006	写x2	显存指针	第一个写指针的高6位 第二个写低8位
	$2007	读写	访问显存数据	指针会在读写后+1或者+32
	$4014	写	DMA访问精灵RAM	通过写一个值$xx, 将CPU内存地址为$xx00-$xxFF的数据复制到精灵内存
*/

func (register *PPURegister) SpriteOverflow() bool {
	return register.ReadBytes(2, 1)[0]<<5>>7 == 0
}

func (register *PPURegister) SpriteHit() bool {
	spriteHit := register.ReadBytes(2, 1)[0]<<6>>7 == 0
	register.ResetVBlank()
	return spriteHit
}

func (register *PPURegister) ResetVBlank() {
	register.WriteBytes(2, register.ReadBytes(2, 1)[0]&191)
}

func (register *PPURegister) VBlank() bool {
	vblankSet := register.ReadBytes(2, 1)[0]<<6>>7 == 0
	register.ResetVBlank()
	return vblankSet
}

func (register *PPURegister) SpriteOAMAddress() byte {
	return register.ReadBytes(3, 1)[0]
}

func (register *PPURegister) SpriteOAMData() byte {
	data := register.ReadBytes(4, 1)[0]
	register.WriteBytes(3, register.SpriteOAMAddress()+1)
	return data
}

func (register *PPURegister) ScrollX() byte {
	return register.ReadBytes(5, 1)[0] >> 4
}
func (register *PPURegister) ScrollY() byte {
	return register.ReadBytes(5, 1)[0] & 0x0F
}
func (register *PPURegister) OAMAddress() byte {
	addr := register.ReadBytes(5, 2)
	return ((addr[0] & 0x3F) << 8) + addr[1]
}

func (register *PPURegister) OAMData() byte {
	data := register.ReadBytes(7, 1)[0] & 0x0F
	//PPU 将会在访问 $2007 后自动增加OAM地址，加1或者32 (基于 $2000 的 D2).
	var offset byte = 1
	if register.OAMMode() == 1 {
		offset = 32
	}
	newOAMAddress := register.OAMAddress() + offset
	register.WriteBytes(5, (newOAMAddress&0x3F)>>8)
	register.WriteBytes(6, newOAMAddress&0x00FF)
	return data
}
