package presentation

import (
	"flag"
	"log"
)

type InputFlag uint8

const (
	All InputFlag = iota
	InProgress
	Open
	Closed
	Init
)

func ParseFlags() InputFlag {
	aFlag := flag.Bool("a", false, "show all todos")
	ipFlag := flag.Bool("*", false, "show in progress todos")
	oFlag := flag.Bool("o", false, "show open todos")
	xFlag := flag.Bool("x", false, "show closed todos")

	flag.Parse()
	flagsCount := flag.NFlag()

	if flagsCount > 1 {
		log.Fatalf("expected 1 argument (a, *, o, x)")
	}
	if flagsCount == 0 {
		if isInitCommand() {
			return Init
		} else {
			return All
		}
	}

	switch {
	case *aFlag:
		return All
	case *ipFlag:
		return InProgress
	case *oFlag:
		return Open
	case *xFlag:
		return Closed
	default:
		return All
	}
}

func isInitCommand() bool {
	commandLineArgs := flag.Args()
	if len(commandLineArgs) == 1 && commandLineArgs[0] == "init" {
		return true
	} else {
		return false
	}
}
