package io

type IORegister struct {
	Memory
}

/**
+---------+----------------------------------------------------------+
    |  $4000  | pAPU Pulse #1 Control Register (W)                       |
    |  $4001  | pAPU Pulse #1 Ramp Control Register (W)                  |
    |  $4002  | pAPU Pulse #1 Fine Tune (FT) Register (W)                |
    |  $4003  | pAPU Pulse #1 Coarse Tune (CT) Register (W)              |
    |  $4004  | pAPU Pulse #2 Control Register (W)                       |
    |  $4005  | pAPU Pulse #2 Ramp Control Register (W)                  |
    |  $4006  | pAPU Pulse #2 Fine Tune Register (W)                     |
    |  $4007  | pAPU Pulse #2 Coarse Tune Register (W)                   |
    |  $4008  | pAPU Triangle Control Register #1 (W)                    |
    |  $4009  | pAPU Triangle Control Register #2 (?)                    |
    |  $400A  | pAPU Triangle Frequency Register #1 (W)                  |
    |  $400B  | pAPU Triangle Frequency Register #2 (W)                  |
    |  $400C  | pAPU Noise Control Register #1 (W)                       |
    |  $400D  | Unused (???)                                             |
    |  $400E  | pAPU Noise Frequency Register #1 (W)                     |
    |  $400F  | pAPU Noise Frequency Register #2 (W)                     |
    |  $4010  | pAPU Delta Modulation Control Register (W)               |
    |  $4011  | pAPU Delta Modulation D/A Register (W)                   |
    |  $4012  | pAPU Delta Modulation Address Register (W)               |
    |  $4013  | pAPU Delta Modulation Data Length Register (W)           |
    +---------+----------------------------------------------------------+
    |  $4014  | Sprite DMA Register (W)                                  |
    |         |                                                          |
    |         |  Transfers 256 bytes of memory into SPR-RAM. The address |
    |         |  read from is $100*N, where N is the value written.      |
    +---------+----------------------------------------------------------+
    |  $4015  | pAPU Sound/Vertical Clock Signal Register (R)            |
    |         |                                                          |
    |         |    D6: Vertical Clock Signal IRQ Availability            |
    |         |           0 = One (1) frame occuring, hence IRQ cannot   |
    |         |               occur                                      |
    |         |           1 = One (1) frame is being interrupted via IRQ |
    |         |    D4: Delta Modulation                                  |
    |         |    D3: Noise                                             |
    |         |    D2: Triangle                                          |
    |         |    D1: Pulse #2                                          |
    |         |    D0: Pulse #1                                          |
    |         |           0 = Not in use                                 |
    |         |           1 = In use                                     |
    |         +----------------------------------------------------------+
    |         | pAPU Channel Control (W)                                 |
    |         |                                                          |
    |         |    D4: Delta Modulation                                  |
    |         |    D3: Noise                                             |
    |         |    D2: Triangle                                          |
    |         |    D1: Pulse #2                                          |
    |         |    D0: Pulse #1                                          |
    |         |           0 = Channel disabled                           |
    |         |           1 = Channel enabled                            |
    +---------+----------------------------------------------------------+
    |  $4016  | Joypad #1 (RW)                                           |
    |         |                                                          |
    |         | READING:                                                 |
    |         |    D4: Zapper Trigger                                    |
    |         |           0 = Pulled                                     |
    |         |           1 = Released (not held)                        |
    |         |    D3: Zapper Sprite Detection                           |
    |         |           0 = Sprite not in position                     |
    |         |           1 = Sprite in front of cross-hair              |
    |         |    D0: Joypad Data                                       |
    |         +----------------------------------------------------------+
    |         | WRITING:                                                 |
    |         | Joypad Strobe (W)                                        |
    |         |                                                          |
    |         |    D0: Joypad Strobe                                     |
    |         |           0 = Clear joypad strobe                        |
    |         |           1 = Reset joypad strobe                        |
    |         +----------------------------------------------------------+
    |         | WRITING:                                                 |
    |         | Expansion Port Latch (W)                                 |
    |         |                                                          |
    |         |    D0: Expansion Port Method                             |
    |         |           0 = Write                                      |
    |         |           1 = Read                                       |
    +---------+----------------------------------------------------------+
    |  $4017  | Joypad #2/SOFTCLK (RW)                                   |
    |         |                                                          |
    |         | READING:                                                 |
    |         |    D7: Vertical Clock Signal (External)                  |
    |         |           0 = Not occuring                               |
    |         |           1 = Occuring                                   |
    |         |    D6: Vertical Clock Signal (Internal)                  |
    |         |           0 = Occuring     (D6 of $4016 affected)        |
    |         |           1 = Not occuring (D6 of $4016 untouchable)     |
    |         |    D4: Zapper Trigger                                    |
    |         |           0 = Pulled                                     |
    |         |           1 = Released (not held)                        |
    |         |    D3: Zapper Sprite Detection                           |
    |         |           0 = Sprite not in position                     |
    |         |           1 = Sprite in front of cross-hair              |
    |         |    D0: Joypad Data                                       |
    |         +----------------------------------------------------------+
    |         | WRITING:                                                 |
    |         | Expansion Port Latch (W)                                 |
    |         |                                                          |
    |         |    D0: Expansion Port Method                             |
    |         |           0 = ???                                        |
    |         |           1 = Read                                       |
    +---------+----------------------------------------------------------+
 */
/**
Transfers 256 bytes of memory into SPR-RAM. The address |
    |         |  read from is $100*N, where N is the value written.
*/
func (register *IORegister) OAMDAM() byte {
		return register.ReadBytes(0x14, 1)[0]
}