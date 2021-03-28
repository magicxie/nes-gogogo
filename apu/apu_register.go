package apu

import "nes6502/io"

type APURegister struct{
	 io.Memory
}
/**
$4000  | pAPU Pulse #1 Control Register (W)                       |
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
    |  $4013  | pAPU Delta Modulation Data Length Register (W)
*/
