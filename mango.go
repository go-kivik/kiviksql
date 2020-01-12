package kiviksql

import (
	"fmt"

	"github.com/pingcap/parser/opcode"
)

// All operators supported by Mango queries.
const (
	And       = "$and"
	Or        = "$or"
	Not       = "$not"
	Nor       = "$nor"
	All       = "$all"
	ElemMatch = "$elemMatch"
	AllMatch  = "$allMatch"
	EQ        = "$eq"
	LT        = "$lt"
	LTE       = "$lte"
	GT        = "$gt"
	GTE       = "$gte"
	NE        = "$ne"
	Exists    = "$exists"
	Type      = "$type"
	In        = "$in"
	Nin       = "$nin"
	Size      = "$size"
	Mod       = "$mod"
	Regex     = "$regex"
)

func mangoOp(op opcode.Op) (string, error) {
	switch op {
	case opcode.LogicAnd:
		return And, nil
	case opcode.LogicOr:
		return Or, nil
	case opcode.Not:
		return Not, nil
	case opcode.EQ:
		return EQ, nil
	case opcode.LT:
		return LT, nil
	case opcode.LE:
		return LTE, nil
	case opcode.GT:
		return GT, nil
	case opcode.GE:
		return GTE, nil
	case opcode.NE:
		return NE, nil
	case opcode.In:
		return In, nil
		// Exists    = "$exists"
		// Type      = "$type"
		// Nin       = "$nin"
		// Size      = "$size"
		// Mod       = "$mod"
		// Regex     = "$regex"
		// Nor       = "$nor"
		// All       = "$all"
		// ElemMatch = "$elemMatch"
		// AllMatch  = "$allMatch"
	}
	return "", fmt.Errorf("unknown operator %s", op)
}
