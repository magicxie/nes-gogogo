package clock

import (
	"fmt"
	"time"
)

type Division struct {
	channel chan int
	cycles  int
}

type Clock struct {
	duration  time.Duration
	Divisions []*Division
}

const (
	HZ  = time.Microsecond
	MHZ = time.Nanosecond
)

func (clock *Clock) SetFrequency(frequency float64, tickUnit time.Duration) {
	clock.duration = tickUnit * time.Duration(1000/frequency)
	fmt.Printf("Duration %d ns, frequency:%f\n", clock.duration, frequency*float64(time.Second.Nanoseconds())/1000)
}

func (clock *Clock) FrequencyDivision(channel chan int, cycles int) {
	clock.Divisions = append(clock.Divisions, &Division{channel, cycles})
}
func (clock *Clock) output() {
	for _,d := range clock.Divisions{
		d.channel <- d.cycles
	}
}
func (clock *Clock) StartTick() {

	//ticker := time.NewTicker(clock.duration)
	var i int64
	for i = 0; i < 5360520; i++ {
		//<-ticker.C
		time.Sleep(clock.duration)
		go clock.output()
		if i == 5360520-1 {
			i = 0
		}
	}
}
