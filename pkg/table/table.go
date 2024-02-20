package table

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
	switch cmpr := md.Comparator.(type) {
	case int:
		switch md.Operator {
		case Equal:
			for _, v := range vals {
				if v.(int) == cmpr {
					md.verified = true
					return nil
				}
			}
		}
	case string:
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
