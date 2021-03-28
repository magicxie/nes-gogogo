package clock

import (
	"testing"
)

func TestClockTick(t *testing.T) {

	var c = Clock{}
	c.SetFrequency(5.37, MHZ)
	var nextTick = make(chan int, 1000000)
	var nextTickB = make(chan int, 1000000)
	c.FrequencyDivision(nextTick, 3)
	c.FrequencyDivision(nextTickB, 9)

	go c.StartTick()
	go func(){
		for i := range nextTick {
			println(i)
		}
	}()
	for i := range nextTickB {
		println(i)
	}

}
