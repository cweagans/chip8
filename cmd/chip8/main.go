package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cweagans/chip8/pkg/cpu"
	"github.com/cweagans/chip8/pkg/graphics"
)

var (
	RomFile      string
	GraphicsMode string
	Debug        bool
	ClockSpeed   int
)

func init() {
	flag.StringVar(&GraphicsMode, "ui", "raylib", "Which UI should the emulator use? Options: raylib (default), termbox.")
	flag.BoolVar(&Debug, "debug", false, "Set debug to true if you want to log CPU internals")
	flag.IntVar(&ClockSpeed, "clock-speed", 60, "Set the CPU clock speed (in Hertz).")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("ROM file argument is required.")
		os.Exit(1)
	}
	RomFile = os.Args[1]
}

func main() {
	// Get a graphics object.
	g := graphics.GetGraphics(GraphicsMode)

	// Load ROM to pass to CPU.
	rom, err := loadRom(RomFile)
	if err != nil {
		fmt.Println("Could not open specified ROM file: " + err.Error())
	}

	// Create a new CPU.
	c := cpu.NewCpu(g, rom, Debug)

	// Set the clock speed based on input.
	c.SetClockSpeed(ClockSpeed)

	// Run the CPU.
	c.Run()
}

func loadRom(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}
