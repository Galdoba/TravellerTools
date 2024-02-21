package table

import "fmt"

const (
	Equal       RelationalOperator = 10
	NotEqual    RelationalOperator = 11
	Less        RelationalOperator = 12
	LessOrEqual RelationalOperator = 13
	More        RelationalOperator = 14
	MoreOrEqual RelationalOperator = 15
)

type Table struct {
	Name      string
	Dicecode  string
	Modifiers []Modifier
}

type TabVal struct {
	Index     int
	Value     string
	NextTable string
	Mods      map[int][]Modifier
}

type Modifier struct {
	Operator   RelationalOperator
	Comparator interface{}
	cmprType   string
	Val        int
	verified   bool
}

func (md *Modifier) Verify(vals ...interface{}) error {
	md.verified = false
	switch cmpr := md.Comparator.(type) {
	default:
		return fmt.Errorf("unknown comparator type %v", cmpr)
	case int:
		switch md.Operator {
		default:
			return fmt.Errorf("unknown operator %v", md.Operator)
		case Equal:
			for _, v := range vals {
				if v.(int) == cmpr {
					md.verified = true
					return nil
				}
			}
		case NotEqual:
			for _, v := range vals {
				if v.(int) != cmpr {
					md.verified = true
					return nil
				}
			}

		case Less:
			for _, v := range vals {
				if v.(int) < cmpr {
					md.verified = true
					return nil
				}
			}
		case LessOrEqual:
			for _, v := range vals {
				if v.(int) <= cmpr {
					md.verified = true
					return nil
				}
			}
		case More:
			for _, v := range vals {
				if v.(int) > cmpr {
					md.verified = true
					return nil
				}
			}
		case MoreOrEqual:
			for _, v := range vals {
				if v.(int) >= cmpr {
					md.verified = true
					return nil
				}
			}
		}
	case string:
		fmt.Errorf("string not implemented")
	}
	return nil
}

func (md *Modifier) Value() int {
	if md.verified {
		return md.Val
	}
	return 0
}

type RelationalOperator int

func (t *Table) SetModifier(n int, is RelationalOperator, comparator interface{}) error {
	mod := newModifier(n, is, comparator)
	t.Modifiers = append(t.Modifiers, mod)
	return nil
}

func newModifier(n int, opr RelationalOperator, cpmr interface{}) Modifier {
	md := Modifier{}
	md.Comparator = cpmr
	md.Operator = opr
	md.Val = n
	return md
}
