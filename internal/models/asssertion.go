package models

const (
	opEqual operator = iota
	opNotEqual
	// opLessThan
	// opGreaterThan
)

type Assertion struct {
	Field    string
	Value    any
	Operator operator
	Result   bool
}

type operator int

func Operator(op operator) string {
	switch op {
	case opEqual:
		return "="
	case opNotEqual:
		return "!="
		// case opLessThan:
		// 	return "<"
		// case opGreaterThan:
		// 	return ">"
	}
	return ""
}

// Assert evaluates the fields of the assertion to true or false
// based on the stated operator
func (a Assertion) Assert(value any) Assertion {
	// there's gotta be a better way
	switch a.Operator {
	case opEqual:
		a.Result = (value == a.Value)
	case opNotEqual:
		a.Result = (value != a.Value)
	}
	return a
}
