package presentation

// TODO: Add commands (init)
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
)

func ParseFlags() InputFlag {
	aFlag := flag.Bool("a", false, "show all todos")
	ipFlag := flag.Bool("*", false, "show in progress todos")
	oFlag := flag.Bool("o", false, "show open todos")
	xFlag := flag.Bool("x", false, "show closed todos")

	flag.Parse()

	if flag.NFlag() > 1 {
		log.Fatalf("expected 1 argument (a, *, o, x)")
	}

	if *aFlag {
		return All
	} else if *ipFlag {
		return InProgress
	} else if *oFlag {
		return Open
	} else if *xFlag {
		return Closed
	} else {
		return All
	}
}
