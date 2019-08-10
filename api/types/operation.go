package types

type Operation uint8

const (
	Unknown Operation = iota
	Add
	Subtract
	Multiply
	Divide
)

var (
	operationStrings = []string{
		"unknown",
		"add",
		"sub",
		"mult",
		"div",
	}

	strToOp = map[string]Operation{
		"add":  Add,
		"sub":  Subtract,
		"mult": Multiply,
		"div":  Divide,
	}
)

func (op Operation) String() string {
	return operationStrings[op]
}

func fromString(s string) Operation {
	op, ok := strToOp[s]
	if !ok {
		return Unknown
	}

	return op
}

func (op *Operation) UnmarshalJSON(bytes []byte) error {
	data := string(bytes)
	if data == "null" {
		return nil
	}

	data = data[1 : len(data)-1]

	*op = fromString(data)

	return nil
}
