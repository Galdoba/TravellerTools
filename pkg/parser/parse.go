package parser

import (
	"fmt"
	"strings"
)

const (
	RuneEOF = 0
)

type ParserFunc = func(*Reader) (*Result, *Error)

type Reader struct {
	data    []byte
	pos     Position
	prevCol int
}

type Position struct {
	offset int
	line   int
	col    int
}

func (p *Position) ToString() string {
	return fmt.Sprintf("offset: %v (ln %v col %v)", p.offset, p.line+1, p.col+1)
}

func NewReader(data string) *Reader {
	r := Reader{}
	r.data = []byte(data)
	return &r
}

func (r *Reader) Read() byte {
	if r.pos.offset >= len(r.data) {
		return RuneEOF
	}
	chr := r.data[r.pos.offset]
	//смещение физической позиции
	r.pos.offset++
	//смещение абстрактной (логической) позиции
	r.prevCol = r.pos.col
	r.pos.col++
	if chr == '\n' {
		r.pos.col = 0
		r.pos.line++
	}

	return chr
}

func (rd *Reader) Data(startPos Position) string {
	return string(rd.data[startPos.offset:rd.pos.offset])

}

func (r *Reader) UnRead() {
	//смещение физической позиции
	r.pos.offset--
	//смещение абстрактной (логической) позиции
	chr := r.data[r.pos.offset]
	r.pos.col--
	if chr == '\n' {
		r.pos.col = r.prevCol
		r.pos.line--
	}
}

func (r *Reader) Save() Position {
	return r.pos
}

func (r *Reader) Restore(p Position) {
	r.pos = p
}

func (r *Reader) Reset() {
	r.pos = Position{}
}

//////////////////////////////////////////////////////////////////////////

type Error struct {
	errors []parseError
}

type parseError struct {
	message string
	pos     Position
}

func NewError() *Error {
	return &Error{}
	//return &Error{[]parseError{parseError{msg, p}}}
}

func (e *Error) Add(p Position, msg string) *Error {
	e.errors = append(e.errors, parseError{msg, p})
	return e
}

func (e *Error) ToString() string {
	str := "Errors:\n"
	l := len(e.errors) - 1
	for i := range e.errors {
		str += fmt.Sprintf("  %v: %v\n", e.errors[l-i].pos.ToString(), e.errors[l-i].message)
	}
	return str
}

////////////////////////////////////////////////////////////////////////
type Result struct {
	kvs []keyval
}

func NewResult(key, val string) *Result {
	return &Result{
		kvs: []keyval{
			{k: key, v: val},
		},
	}
}

func (res *Result) ToString() string {
	if res == nil {
		return "ok"
	}
	str := ""
	for _, kv := range res.kvs {
		str += fmt.Sprintf("key: %v  val: %v\n", kv.k, kv.v)
	}
	return str
}

func AppendResult(res, new *Result) *Result {
	if new == nil {
		return res
	}
	if res == nil {
		res = &Result{}
	}
	res.kvs = append(res.kvs, new.kvs...)
	return res
}

type keyval struct {
	k string
	v string
}

////////////////////////////////////////////////////////////////////////

func Seq(args ...interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		res := (*Result)(nil)
		for _, arg := range args {
			pos := rd.Save()
			//fmt.Printf("-%T*\n", arg)
			switch val := arg.(type) {
			default:
				panic(fmt.Sprintf("unallowed state %T", val))
			case int32:
				chr := rd.Read()
				v := byte(val)
				if chr != v {
					return nil, NewError().Add(pos, fmt.Sprintf("expected %q, got %q", v, chr))
				}
			case string:
				if val == "" {
					continue
				}
				for _, s := range []byte(val) {
					chr := rd.Read()
					if chr != s {
						return nil, NewError().Add(pos, fmt.Sprintf("read string: %v: expected %q, got %q", val, s, chr))
					}
				}
			case ParserFunc:
				rs, err := val(rd)
				if err != nil {
					return nil, err
				}
				res = AppendResult(res, rs)
			} //switch
		} //for
		return res, nil
	} //func
}

//подумать над именем //choice?
func Choose(args ...interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		res := (*Result)(nil)
		pos := rd.Save()
	argLoop:
		for _, arg := range args {
			//fmt.Printf("-%T*\n", arg)
			rd.Restore(pos)
			switch val := arg.(type) {
			default:
				panic(fmt.Sprintf("unallowed state %T", val))
			case int32:
				chr := rd.Read()
				v := byte(val)
				if chr == v {
					return res, nil
				}
			case string:
				if val == "" {
					return res, nil
				}
				for _, s := range []byte(val) {
					chr := rd.Read()
					if chr != s {
						continue argLoop
					}
				}
				return res, nil
			case ParserFunc:
				rs, err := val(rd)
				if err == nil {
					res = AppendResult(res, rs)
					return res, nil
				}
			} //switch
		} //for
		return nil, NewError().Add(pos, "choice invalid")
	} //func
}

//Optional 0 или 1 arg
func Optional(arg interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		fn := Choose(arg, "")
		res, _ := fn(rd)
		return res, nil
	}
}

//abcd
/*
a != c
keep a
b != c
keep b
c == c
return ab
*/

func ZeroOrMany(arg interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		fn := Seq(arg)
		res := (*Result)(nil)
		for {
			pos := rd.Save()
			resNew, err := fn(rd)
			res = AppendResult(res, resNew)
			if err != nil {
				rd.Restore(pos)
				return res, nil
			}
		}
	}
}

//TODO: разобраться с тестами и понять работает ли оно
//пока фиксированное колличество символов
func Not(args ...interface{}) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		res := (*Result)(nil)
		for _, arg := range args {
			pos := rd.Save()
			//fmt.Printf("-%T*\n", arg)
			switch val := arg.(type) {
			default:
				panic(fmt.Sprintf("unallowed state %T", val))
			case int32:
				chr := rd.Read()
				v := byte(val)
				same := true
				if chr != v {
					same = false
				}
				if !same {
					continue
				}
				return nil, NewError().Add(pos, fmt.Sprintf("expected difference: %q, got %q", v, chr))
			case string:
				if val == "" {
					continue
				}
				same := true
				for _, s := range []byte(val) {
					chr := rd.Read()
					if chr != s {
						same = false
					}
				}
				if !same {
					continue
				}
				return nil, NewError().Add(pos, fmt.Sprintf("read string: %v: expected difference", val))
			case ParserFunc:
				//fmt.Println("case ParserFunc:")
				rs, err := val(rd)
				//fmt.Println("rs, err:", rs, err)
				if err == nil {
					//panic(4)
					fn2 := Choose("", arg)
					res, _ := fn2(rd)
					//fmt.Println(res)
					return res, nil
					return Choose("", arg)(rd)
					return nil, NewError().Add(pos, fmt.Sprintf("read string: %v: expected difference func", arg))
				} else {
					return nil, NewError().Add(pos, fmt.Sprintf("read string: %v: expected difference func", arg))
				}
				res = AppendResult(res, rs)
			} //switch
		} //for
		return res, nil
	} //func
}

//////////////////////////////////////////////////////////////////////////
//-Keep сохраняет найденное
func Keep(name string, arg interface{}) ParserFunc {
	//res, err := Choose(arg, "")
	return func(rd *Reader) (*Result, *Error) {
		pos := rd.Save()
		fn := Seq(arg)
		res, err := fn(rd)
		if err != nil {
			return nil, err.Add(pos, fmt.Sprintf("Keep %q ненашел то что искал ", name))
		}
		res = AppendResult(res, NewResult(name, rd.Data(pos)))
		return res, nil
	}

}

//////////////////////////////////////////////////////////////////////////

func Func(fn func(chr byte) bool) ParserFunc {
	return func(rd *Reader) (*Result, *Error) {
		//сравнивается chr и чтение их ридера по chr из вне
		if fn(rd.Read()) {
			return nil, nil
		}
		rd.UnRead()
		return nil, NewError()
	}
}

func Ident() ParserFunc {
	return Seq(Func(Alpha), ZeroOrMany(Choose(Func(Alpha), Func(Digit))))
}

func Alpha(chr byte) bool {
	return LowerAlpha(chr) || UpperAlpha(chr)
}

func LowerAlpha(chr byte) bool {
	if chr >= 'a' && chr <= 'z' {
		return true
	}

	return false
}

func UpperAlpha(chr byte) bool {
	if chr >= 'A' && chr <= 'Z' {
		return true
	}
	return false
}

func Digit(chr byte) bool {
	if chr >= '0' && chr <= '9' {
		return true
	}
	return false
}

func HashAlpha(chr byte) bool {
	if chr >= '0' && chr <= '9' && chr >= 'A' && chr <= 'Z' && chr != 'I' && chr != 'O' {
		return true
	}
	return false
}

//примерсоздания своих алфавитов
func EHex(chr byte) bool {
	const alph = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	return strings.Contains(alph, string(chr))
}

func Space(chr byte) bool {
	const alph = " 	\r\n\t"
	return strings.Contains(alph, string(chr))
}

func UniversalProfile() ParserFunc {
	return Seq(Func(EHex), Func(EHex), Func(EHex), Func(EHex), Func(EHex), Func(EHex), Func(EHex), '-', Func(EHex))
}

/////////////////////////////////////////////////////

func Run(data string, fn ParserFunc) (*Result, error) {
	rd := NewReader(data)
	res, e := fn(rd)
	if e != nil {
		return nil, fmt.Errorf("run(): %v", e.ToString())
	}
	return res, nil
}
