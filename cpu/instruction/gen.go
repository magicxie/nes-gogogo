package instruction

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type AD struct {
	add   string
	oper  string
	desc  string
	bytes string
	cycle string
}
type GT struct {
	name    string
	desc    string
	address []AD
}

func Gen() {
	file, err := os.Open("./1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := []GT{}
	idx := 0

	for scanner.Scan() {
		v := scanner.Text()

		if len(v) == 3 {
			idx = 0
			gt := GT{}
			gt.name = v
			gt.address = make([]AD, 0)
			arr = append(arr, gt)
			idx++
		} else {
			ele := &arr[len(arr)-1]
			if idx == 1 {
				ele.desc = v
				idx = 2
			}
			if idx == 3 {
				ap := strings.Split(v, "\t")
				if len(ap) != 5 {
					fmt.Printf("%v", ap)
				}

				e := AD{
					ap[0],
					ap[2],
					ap[1],
					ap[3],
					ap[4],
				}
				ele.address = append(ele.address, e)
			}
			if v == "addressing\tassembler\topc\tbytes\tcyles" {
				idx = 3
			}

		}

	}

	for _, s := range arr {
		fmt.Printf("%s = MicroInstruction{\n\t\t\"%s\",\n\t\t\"%s\",\n\t\tfunc(operand []byte, bus Bus)  {},\n\t}\n", s.name, s.name, s.desc)
	}

	for _, s := range arr {
		for _, i := range s.address {

			var addd = strings.ToLower(i.add)
			addd = strings.Title(addd)
			addd = strings.ReplaceAll(addd, ",", "")
			addd = strings.ReplaceAll(addd, ")", "")
			addd = strings.ReplaceAll(addd, "(", "")

			fmt.Printf("0x%s: Instruction{\n\t\t%s, %s, \"%s\",%s, addressing.%s{},\n\t},", i.oper, i.bytes, i.cycle,  i.desc,s.name, addd)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
