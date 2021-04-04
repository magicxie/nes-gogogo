package rom

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

type HEX struct {
	str string
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}
func (color RGB) rgb2hex() HEX {
	r := t2x(color.red)
	g := t2x(color.green)
	b := t2x(color.blue)
	return HEX{r + g + b}
}

func TestResolver_Resolve(t *testing.T) {

	fs, _ := ioutil.ReadFile("/Users/eggcanfly/go/src/nes-test-roms/spritecans-2011/spritecans.nes")

	r := &Resolver{}
	rom := r.Resolve(fs)
	rom.Summary()

	//调色板内存位于0x3F00-0x3F1F（32字节），由两个16字节的调色板组成
	//一个用于背景调色（0x3F00），一个用于精灵调色（0x3F10）。
	//其中每个字节表示一个0-63的数，索引到系统调色板，下面看到的颜色都来自系统调色板

	if rom.Header.CHRMirrors > 0 {

		fmt.Printf("Pattern Table\n")
		patternTable0 := rom.Character.data[0:0x0FFF]
		patternTable1 := rom.Character.data[0x1000:0x1FFF]

		for i := 0; i < len(patternTable0); i++ {

			if i%8 == 0 {
				fmt.Printf("$%04X\n", i)
			}

			a := patternTable0[i]
			b := patternTable1[i]
			as := fmt.Sprintf("%08b", a)
			bs := fmt.Sprintf("%08b", b)

			//image := &Image{}
			file, _ := os.Create("img/" + strconv.Itoa(i) + ".jpeg")
			rgba := image.NewRGBA(image.Rect(0, 0, 64, 64))
			defer file.Close()
			println(rgba)

			for idx, e := range strings.Split(as, "") {
				abit, _ := strconv.Atoi(e)
				bbit, _ := strconv.Atoi(strings.Split(bs, "")[idx])
				fmt.Printf("%d\t", abit+bbit*2)
				//rgba.SetRGBA()
			}
			println("")

		}
	}

	type NameTable struct {
		nameTable      []byte
		attributeTable []byte
	}
	nameTable := []NameTable{}
	//FC有4个名称表，位于0x2000-0x2FFF，一共4KB，每个名称表1024字节。其中前960字节存储实际名称表，后64字节存储属性表。
	for ii := 0; ii <= 3; ii++ {
		start := 0x2000 + ii*0x1000

		physicalNameTable := rom.Character.data[start : start+960]
		attributeTable := rom.Character.data[start+961 : start+0x1000]

		nameTable = append(nameTable, NameTable{physicalNameTable, attributeTable})

		println(len(nameTable), len(attributeTable))

		for _, c := range nameTable {
			fmt.Printf("%X", c)
		}
		for _, c := range attributeTable {
			fmt.Printf("%X", c)
		}
	}

}
