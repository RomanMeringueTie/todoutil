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

var symbolStatus = map[rune]Status{
	'o': Open,
	'x': Closed,
	'*': InProgress,
}

func (status Status) String() string {
	return statusName[status]
}

func SymbolToStatus(symbol rune) Status {
	return SymbolToStatus(symbol)
}
