package parser

import (
	"strings"
	"testing"
)

//|Drinax|2223|A43645A-E|714|||NaHu|M1 V|K|{ +1 }|1|(B34+3)|[657G]|B|9|396|10|5|-107|-17|Ni||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated
/*
SEP	IDENT(ZERO+(IDENT))	SEP SEQ(DIGIT,DIGIT,DIGIT,DIGIT) SEP

*/
var inputCorrect = []struct {
	in string
	fn ParserFunc
}{
	{in: "abcdef", fn: Seq("abcdef")},
	//{in: "abdd", fn: Not("abcd")},
	//{in: "a", fn: Not('b')},
	//{in: "hb", fn: Seq(Not("a"), "b")},
	{in: "b", fn: Not("a")},

	{in: "", fn: Not("a")},
	//{in: "abc", fn: Not(Seq("a", "bc"))},    {sdkfbgs}   Seq('{', ZeroOrMany(Not('}')) ..., '}')
	//{in: "abc", fn: Not(Seq("abc"))},
	{in: "abcdef", fn: Seq("abc", 'd', "ef")},
	{in: "abcf2", fn: Seq("abc", Not(Seq("vs")), "f2")},
	{in: "abcdef", fn: Seq("abc", Seq("de"), "f")},
	{in: "ab", fn: Choose("ab", 'b', Seq("c"))},
	{in: "b", fn: Choose("ab", 'b', Seq("c"))},
	{in: "c", fn: Choose("ab", 'b', Seq("c"))},
	{in: "aa\tbb", fn: Seq("aa", Seq(Func(Space)), "bb")},
	{in: "asd A456789-D", fn: Optional(Seq(ZeroOrMany(Ident()), UniversalProfile()))},
	//|Drinax|2223|
	{in: "|Drinax|2223|A43645A-E", fn: Seq('|', Seq(ZeroOrMany(Ident())), "|", Seq(Func(Digit), Func(Digit), Func(Digit), Func(Digit)), '|', Seq(UniversalProfile()))},
}

var inputCorrectKeep = []struct {
	in  string
	fn  ParserFunc
	out []string
}{

	{in: "abcdef", fn: Keep("val", "abcdef"), out: []string{"val:abcdef"}},
	{
		in:  "key:a555",
		fn:  Seq("key:", Keep("keyFound", Ident())),
		out: []string{"keyFound:a555"},
	},

	{
		in: "aaa:b555",
		fn: Optional(Seq(
			Keep("key", Ident()),
			':',
			Keep("val", Ident()),
		)), //Seq("key:", Keep("keyFound", Ident())),
		out: []string{"key:aaa", "val:b555"},
	},
	{
		in: "|Drinax|2223|",
		fn: Seq(
			'|',
			Keep("name", Seq(ZeroOrMany(Ident()))),
			"|",
			Keep("hex", Seq(Func(Digit), Func(Digit), Func(Digit), Func(Digit))),
			'|',
		),
		out: []string{"name:Drinax", "hex:2223"},
	},
}

var inputIncorrect = []struct {
	in string
	fn ParserFunc
}{
	{in: "", fn: Seq("abcdef")},
	//{in: "abc", fn: Not(Seq("ab", "cc"))},
	{in: "abcdef", fn: Seq("abc", 't', "ef")},
	{in: "abcdef", fn: Seq("abc", Seq("ge"), "f")},
	{in: "", fn: Choose("ab", 'b', Seq("c"))},
	{in: "d", fn: Choose("ab", 'b', Seq("c"))},
	{in: "d", fn: Choose("ab", 'b', Seq("c"))},
	{in: "a", fn: Not("a")},
}

func TestCorrectParser(t *testing.T) {
	t.Logf("START TestCorrect")
	for i, input := range inputCorrect {
		t.Logf("----------------------")
		// rd := NewReader(input.in)
		// r, e := input.fn(rd)
		r, e := Run(input.in, input.fn)
		if e != nil {
			t.Errorf("%v: %q", i, input.in)
			t.Errorf("    %v", e)
		} else {
			t.Logf("%v: %q", i, input.in)
			t.Logf("    %v", r.ToString())
		}

	}

}

func TestCorrectKeepParser(t *testing.T) {
	t.Logf("START TestCorrectKeep")
	for i, input := range inputCorrectKeep {
		t.Logf("----------------------TestCorrectKeep-------------")
		out := formatOut(input.out)
		//rd := NewReader(input.in)
		//r, e := input.fn(rd)
		r, e := Run(input.in, input.fn)
		if e != nil {
			t.Errorf("%v: %q", i, input.in)
			t.Errorf("    %v", e)
		} else {
			t.Logf("%v: %q", i, input.in)
			t.Logf("    %v", r.ToString())
		}
		if !checkResult(t, r, out) {
			t.Errorf("результат не совпал:\n")
			t.Errorf("  got = %v\n", r)
			t.Errorf("  exp = %v\n", detectedToString(out))
		}
	}

}

type detected struct {
	key   string
	val   string
	found bool
}

func detectedToString(det []*detected) []detected {
	//return fmt.Sprintf()
	if len(det) == 0 {
		return nil
	}
	detr := []detected{}
	for _, d := range det {
		detr = append(detr, *d)
	}
	return detr
}

func formatOut(out []string) []*detected {
	det := []*detected{}
	for _, data := range out {
		dt := strings.Split(data, ":")
		det = append(det, &detected{key: dt[0], val: dt[1]})
	}
	return det
}

func checkResult(t *testing.T, got *Result, expected []*detected) bool {
	for _, res := range got.kvs {
		key := res.k
		val := res.v
		for _, exp := range expected {
			//t.Errorf("--key:'%v'", exp.key)
			if key == exp.key {
				if val == exp.val {
					if exp.found == false {
						exp.found = true
						break
					}
				}
			}

		}

	}
	for _, ex := range expected {
		if !ex.found {
			return false
		}
	}
	return true
}

func TestInCorrectParser(t *testing.T) {

	t.Logf("START TestInCorrect")
	for i, input := range inputIncorrect {
		t.Logf("----------------------")
		//rd := NewReader(input.in)
		//r, e := input.fn(rd)
		r, e := Run(input.in, input.fn)
		if e != nil {
			t.Logf("%v: %q", i, input.in)
			t.Logf("    %v", e)
		} else {
			t.Errorf("%v: %q", i, input.in)
			t.Errorf("    %v", r.ToString())
		}
	}

}
