package main

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/cli/command"
)

func Test_Traffic(t *testing.T) {
	fmt.Println("this is a test!")
	rn, err := command.New(
		command.CommandLineArguments("go run main.go traffic --worldname Drinax"),
		command.Set(command.TERMINAL_ON),
	)
	if err != nil {
		t.Errorf("internal error: %v", err.Error())
	}
	err = rn.Run()
	if err != nil {
		t.Errorf("execution error: %v", err.Error())
	}
}
