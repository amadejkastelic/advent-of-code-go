package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/amadejkastelic/advent-of-code-go/internal/utils"
)

var (
	initPtr   = flag.Bool("init", false, "Initialize solution template")
	yearPtr   = flag.Int("year", 2025, "Year of AoC puzzle")
	dayPtr    = flag.Int("day", 1, "Day of AoC puzzle")
	simplePtr = flag.Bool("simple", false, "Fetch simple input")
)

const pathTemplate = "./internal/solutions/%d/day%d/solution.go"

func main() {
	flag.Parse()

	if *initPtr {
		if err := initSolutionTemplate(*yearPtr, *dayPtr); err != nil {
			panic(err)
		}
		fmt.Println("Initialized solution template: " + fmt.Sprintf(pathTemplate, *yearPtr, *dayPtr))
		return
	}

	inputPath, err := utils.FetchInput(*yearPtr, *dayPtr, *simplePtr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Input path:", inputPath)

	solutionPath := fmt.Sprintf(pathTemplate, *yearPtr, *dayPtr)
	fmt.Println("Solution path:", solutionPath)

	if fileInfo, err := os.Stat(solutionPath); err != nil || fileInfo.IsDir() {
		panic("Solution file does not exist")
	}

	cmd := exec.Command("go", "run", solutionPath, "-input", inputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func initSolutionTemplate(year int, day int) error {
	solutionDir := fmt.Sprintf("./internal/solutions/%d/day%d", year, day)
	if err := os.MkdirAll(solutionDir, os.ModePerm); err != nil {
		return err
	}

	solutionPath := fmt.Sprintf("%s/solution.go", solutionDir)
	templatePath := "./templates/solution.go"

	input, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	if err := os.WriteFile(solutionPath, input, 0644); err != nil {
		return err
	}

	return nil
}
