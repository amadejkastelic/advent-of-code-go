package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/draffensperger/golp"

	"github.com/amadejkastelic/advent-of-code-go/internal/sliceutils"
	"github.com/amadejkastelic/advent-of-code-go/internal/utils"
)

const maxPresses = 1000

var (
	inputPath = flag.String("input", "", "Path to input file")
	part      = flag.Int("part", 0, "Part of the puzzle to solve (1 or 2). If 0, solve both parts.")
)

func parseInput(s string) []*Machine {
	machines := make([]*Machine, 0)
	for line := range strings.SplitSeq(strings.Trim(s, "\n"), "\n") {
		machines = append(machines, MachineFromString(line))
	}
	return machines
}

func main() {
	flag.Parse()

	if *inputPath == "" {
		panic("Input file path is required")
	}

	puzzle, err := os.ReadFile(*inputPath)
	if err != nil {
		panic(err)
	}

	machines := parseInput(string(puzzle))

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(machines), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(machines), time.Since(start))
	}
}

type Button struct {
	ToggleIndexes []int
}

type Machine struct {
	buttons []*Button

	numBits      int
	initialState utils.Binary
	state        utils.Binary
	desiredState utils.Binary

	desiredJoltage []int
	joltage        []int
}

func MachineFromString(s string) *Machine {
	m := &Machine{
		state:        0b0,
		desiredState: 0b0,
	}
	for f := range strings.FieldsSeq(s) {
		switch f[0] {
		case '[':
			for _, c := range f[1 : len(f)-1] {
				m.desiredState = utils.ShiftLeft(m.desiredState, 1)
				if c == '#' {
					m.desiredState = utils.SetLSB(m.desiredState)
				} else {
					m.desiredState = utils.ClearLSB(m.desiredState)
				}
				m.state = utils.ShiftLeft(m.state, 1)
				m.numBits++
			}
			m.initialState = m.state
		case '(':
			b := &Button{}
			for n := range strings.SplitSeq(f[1:len(f)-1], ",") {
				b.ToggleIndexes = append(b.ToggleIndexes, utils.ValOrPanic(strconv.Atoi(n)))
			}
			m.buttons = append(m.buttons, b)
		case '{':
			for j := range strings.SplitSeq(f[1:len(f)-1], ",") {
				m.desiredJoltage = append(m.desiredJoltage, utils.ValOrPanic(strconv.Atoi(j)))
			}
			m.joltage = make([]int, len(m.desiredJoltage))
		}
	}
	return m
}

func (m *Machine) Reset() {
	m.state = m.initialState
	m.joltage = make([]int, len(m.desiredJoltage))
}

func (m *Machine) PressButton(b *Button) {
	for _, toggleIndex := range b.ToggleIndexes {
		m.state = utils.Toggle(m.state, m.numBits-1-toggleIndex)
		m.joltage[toggleIndex] += 1
	}
}

func (m *Machine) NumButtonPressesToStart() int {
	if m.desiredState == m.state {
		return 0
	}

	for presses := 1; presses <= maxPresses; presses++ {
		for _, combination := range sliceutils.CombinationsWithReplacement(m.buttons, presses) {
			for _, btn := range combination {
				m.PressButton(btn)
			}

			if m.desiredState == m.state {
				return presses
			}
			m.Reset()
		}
	}

	return 0
}

func (m *Machine) NumButtonPressesToReachJoltage() int {
	if slices.Equal(m.desiredJoltage, m.joltage) {
		return 0
	}

	numButtons := len(m.buttons)
	numJoltages := len(m.desiredJoltage)

	lp := golp.NewLP(0, numButtons)
	lp.SetVerboseLevel(golp.NEUTRAL)

	objectiveCoeffs := make([]float64, numButtons)
	for i := range numButtons {
		objectiveCoeffs[i] = 1.0
	}
	lp.SetObjFn(objectiveCoeffs)

	for i := range numButtons {
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, float64(maxPresses))
	}

	for i := 0; i < numJoltages; i++ {
		var entries []golp.Entry
		for j, btn := range m.buttons {
			if slices.Contains(btn.ToggleIndexes, i) {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		targetValue := float64(m.desiredJoltage[i])
		if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
			panic(err)
		}
	}

	status := lp.Solve()

	if status != golp.OPTIMAL {
		return 0
	}

	solution := lp.Variables()
	totalPresses := 0
	for _, val := range solution {
		totalPresses += int(val + 0.5)
	}

	return totalPresses
}

func solvePart1(machines []*Machine) any {
	res := 0

	for _, m := range machines {
		res += m.NumButtonPressesToStart()
	}

	return res
}

func solvePart2(machines []*Machine) any {
	res := 0

	for _, m := range machines {
		res += m.NumButtonPressesToReachJoltage()
	}

	return res
}
