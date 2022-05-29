package day8

import (
	"strconv"
	"strings"
)

// Console describes the game console.
type Console struct {
	Seen         map[int]bool
	Instructions []Instruction
	Loop         bool
	Accumulator  int
}

// Instruction is a tuple of a command and integer value
type Instruction struct {
	Command string
	Value   int
}

// Parse takes a slice of strings and interprets them into instructions
// for the Console.
func (c *Console) Parse(input []string) {
	for _, line := range input {
		cmdValue := strings.Split(line, " ")
		value, _ := strconv.Atoi(cmdValue[1])
		c.Instructions = append(c.Instructions, Instruction{Command: cmdValue[0], Value: value})
	}
}

// Run goes through the instructions and halts on detecting a loop.
func (c *Console) Run(kind string) {
	switch kind {
	case "loop":
		c.Runloop()
	case "fixloop":
		c.Fixloop()
	}
}

// Runloop runs the console with the loop.
func (c *Console) Runloop() {
	pos := 0
	for !c.Loop && pos < len(c.Instructions) {
		if c.Seen[pos] {
			c.Loop = true
			break
		}
		cmd := c.Instructions[pos]
		c.Seen[pos] = true
		switch cmd.Command {
		case "nop":
			pos++
		case "acc":
			c.Accumulator += cmd.Value
			pos++
		case "jmp":
			pos += cmd.Value
		}
	}
}

// Fixloop runs the console with the loop fixed.
func (c *Console) Fixloop() {
	for i, instruction := range c.Instructions {
		instructions := []Instruction{}
		instructions = append(instructions, c.Instructions[:i]...)
		if instruction.Command == "nop" {
			instructions = append(instructions, Instruction{Command: "jmp", Value: instruction.Value})
		} else if instruction.Command == "jmp" {
			instructions = append(instructions, Instruction{Command: "nop", Value: instruction.Value})
		} else {
			instructions = append(instructions, Instruction{Command: instruction.Command, Value: instruction.Value})
		}
		instructions = append(instructions, c.Instructions[i+1:]...)
		seen := make(map[int]bool)
		d := &Console{Instructions: instructions, Seen: seen}
		d.Runloop()
		if !d.Loop {
			c.Accumulator = d.Accumulator
			break
		}
	}
}

// Output emits a value of the game console of a certain kind (like accumulator).
func (c *Console) Output() int {
	return c.Accumulator
}

// NewConsole returns a Console ready to parse input and run.
func NewConsole() *Console {
	seen := make(map[int]bool)
	return &Console{Seen: seen}
}

// RunConsole runs the game console and returns a value of kind
// just before an endless loop is detected.
func RunConsole(input []string, kind string) int {
	console := NewConsole()
	console.Parse(input)
	console.Run(kind)
	return console.Output()
}
