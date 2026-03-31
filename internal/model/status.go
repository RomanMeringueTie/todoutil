package model

type Status uint8

const (
	Open Status = iota
	Closed
	InProgress
)

var statusName = map[Status]string{
	Open:       "Open",
	Closed:     "Closed",
	InProgress: "InProgress",
}

var symbolStatus = map[string]Status{
	"[o]": Open,
	"[x]": Closed,
	"[*]": InProgress,
}

func (status Status) String() string {
	return statusName[status]
}

func SymbolToStatus(symbol string) (Status, bool) {
	status, isFind := symbolStatus[symbol]
	if isFind {
		return status, true
	} else {
		return Open, false
	}
}
