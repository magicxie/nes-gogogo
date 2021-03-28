package bus

const (
	WRITE = 1
	READ  = 0
)

type Payload struct {
	address uint16
	signal  byte
	bytes   int
}

type Bus struct {
	payload     chan Payload
	dataBus     chan []byte
	mapper Mapper
	//Ram         *Ram
	//IORegisters *IORegisters
}

func (bus *Bus) Init(mapper Mapper) {
	bus.payload = make(chan Payload, 1)
	bus.dataBus = make(chan []byte, 2)
	bus.mapper = mapper
}

func (bus *Bus) address(bytes int) {
	payload := <-bus.payload

	if payload.signal == READ {
		bus.dataBus <- bus.mapper.Read(payload.address, bytes)//bus.Ram.ReadBytes(payload.address, bytes)
	}

	if payload.signal == WRITE {
		bs := <-bus.dataBus
		bus.mapper.Write(payload.address, bs)
	}
}

func (bus *Bus) ReadByte(address uint16) chan []byte {
	bus.Send(READ, address, 1)
	return bus.dataBus
}

func (bus *Bus) ReadWord(address uint16) chan []byte {
	bus.Send(READ, address, 2)
	return bus.dataBus
}

func (bus *Bus) Read(address uint16, bytes int) chan []byte {
	bus.Send(READ, address, bytes)
	return bus.dataBus
}

func (bus *Bus) Write(address uint16, data []byte) {
	bus.Send(WRITE, address, len(data))
	bus.dataBus <- data
}

func (bus *Bus) WriteByte(address uint16, data []byte) {
	bus.dataBus <- data[0:1]
	bus.Send(WRITE, address, 1)
}

func (bus *Bus) WriteWord(address uint16, data []byte) {
	bus.dataBus <- data[0:2]
	bus.Send(WRITE, address, 2)
}

func (bus *Bus) Send(signal byte, address uint16, bytes int) {
	payload := Payload{address, signal, bytes}
	bus.payload <- payload
	bus.address(bytes)
}

type Mapper interface {
	Read(address uint16, bytes int) []byte
	Write(address uint16, data []byte)
}