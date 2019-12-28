package main

import "fmt"

type CPU struct{}
type Memory struct{}
type SolidStateDrive struct{}

func (cpu CPU) freeze() {
	fmt.Println("Freezing processor")
}

func (cpu CPU) jump(position string) {
	fmt.Println("Jumping to:", position)
}

func (cpu CPU) execute() {
	fmt.Println("Executing")
}

func (memory Memory) load(position, data string) {
	fmt.Println("Loading from " + position + " data " + data)
}

func (ssd SolidStateDrive) read(lba, size string) string {
	return "Some data from sector " + lba + " with size " + size
}

type ComputerFacade struct {
	cpu    CPU
	memory Memory
	ssd    SolidStateDrive
}

func NewComputerFacade() ComputerFacade {
	computer := ComputerFacade{
		CPU{},
		Memory{},
		SolidStateDrive{},
	}
	return computer
}

func (computer ComputerFacade) start() {
	computer.cpu.freeze()
	computer.memory.load("0x00", computer.ssd.read("100", "1024"))
	computer.cpu.jump("0x00")
	computer.cpu.execute()
}

func main() {
	computer := NewComputerFacade()
	computer.start()

	// Freezing processor
	// Loading from 0x00 data Some data from sector 100 with size 1024
	// Jumping to: 0x00
	// Executing
}
