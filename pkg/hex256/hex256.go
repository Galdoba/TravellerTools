package hex256

import (
	"fmt"
	"strings"
)

var byteToString map[byte]string
var stringToByte map[string]byte
var glyphs []string

func init() {
	byteToString = make(map[byte]string)
	stringToByte = make(map[string]byte)
	glyphs = []string{
		//	/* */ "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F",

		/*0        */ "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F",

		/*1	   */ "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X",

		/*2	   */ "Y", "Z", "!", "#", "$", "?", "@", "*", "€", "†", "‡", "ƿ", "ǀ", "ǁ", "Δ", "Λ",

		/*3	   */ "Ξ", "Π", "Σ", "Φ", "Ψ", "Ω", "ϐ", "ϑ", "Ϛ", "Ϟ", "Ϡ", "Б", "Д", "Ж", "Л", "Ц",

		/*4	   */ "Ч", "Ш", "Ъ", "Ы", "Э", "Ю", "Я", "Ѡ", "Ѣ", "Ѧ", "Ѫ", "Ѯ", "Ѳ", "҂", "Ҕ", "Բ",

		/*5	   */ "Գ", "Է", "Ը", "Թ", "Ժ", "Ի", "Ծ", "Հ", "Ղ", "Ո", "Պ", "Վ", "Ր", "Ւ", "۵", "༡",

		/*6	   */ "₡", "₢", "₣", "₤", "₥", "₦", "₧", "₨", "₩", "₪", "₫", "₭", "₮", "←", "↑", "→",

		/*7	   */ "↓", "↔", "↕", "∃", "∅", "∇", "∈", "⊔", "⊕", "⊖", "⊗", "⊘", "⊙", "⊛", "⊜", "⊞",

		/*8	   */ "⊟", "⊠", "⊡", "⍊", "⍏", "⍑", "⍙", "⍚", "⍜", "⍠", "⍡", "⍢", "⍲", "⍺", "␢", "◉",

		/*9        */ "⸢", "⸣", "⸤", "⸥", "⟂", "⟅", "⟆", "⟦", "⟧", "⟨", "⟩", "⟪", "⟫", "¥", "⌿", "⍀",

		/*A        */ "⍃", "⍄", "⍅", "⍆", "⍇", "⍈", "⍉", "⍊", "⍋", "⍌", "⍍", "⍎", "⍏", "⍐", "⍑", "⍒",

		/*B        */ "⌫", "⌸", "⌹", "⌺", "⌻", "⌼", "⌽", "⌾", "⍟", "⍠", "⍡", "⍢", "⍭", "⍮", "⍯", "⍰",

		/*C        */ "☠", "☡", "☢", "☣", "☤", "☥", "☦", "☧", "☨", "☩", "☪", "☫", "☬", "☭", "☮", "☯",

		/*D        */ "☰", "☸", "☹", "☺", "☻", "☼", "☽", "☾", "☿", "♀", "♁", "♂", "♃", "♄", "♅", "♆",

		/*E        */ "♇", "♔", "♕", "♖", "♗", "♘", "♙", "♚", "♛", "♜", "♝", "♞", "♟", "♠", "♡", "♢",

		/*F        */ "♣", "♤", "♥", "♦", "♧", "♨", "♩", "♪", "♫", "♬", "♭", "♮", "♯", "♰", "♱", "♲",
	}
	for i := 0; i < 256; i++ {
		bt := byte(i)
		byteToString[bt] = glyphs[i]
		stringToByte[glyphs[i]] = bt
	}
}

func ByteOf(s string) byte {
	return stringToByte[s]
}

type hx256 struct {
	raw           byte
	glyph         string
	index         int
	subIndex      int
	data          string
	reformedOrder []byte
}

type Ehex interface {
	Value() int
	String() string
	Data() string
	Set(interface{}) error
	Add(int) error
}

func New(input ...interface{}) (*hx256, error) {
	eh := hx256{}
	for _, inp := range input {
		if err := eh.Set(inp); err != nil {
			return nil, err
		}
	}
	return &eh, nil
}

func (eh *hx256) Set(input interface{}) error {
	switch v := input.(type) {
	default:
		return fmt.Errorf("unsupported type passed '%v'", v)
	case byte:
		eh.setByte(v)
	case int:
		switch len(eh.reformedOrder) {
		case 0:
			if v < 0 || v > 255 {
				return fmt.Errorf("int value (%v) incorrect", v)
			}
			bt := byte(v)
			eh.setByte(bt)
		default:
			for i, bt := range eh.reformedOrder {
				if i != v {
					continue
				}
				eh.setByte(bt)
				break
			}
			return fmt.Errorf("int value (%v) not allowed", v)
		}
	case string:
		sl := strings.Split(v, "")
		switch len(sl) {
		default:
			if eh.glyph == "" {
				return fmt.Errorf("can't set description '%v' to unasigned value", v)
			}
			eh.data = v
		case 1:
			bt := stringToByte[v]
			eh.setByte(bt)
		}
	case []byte:
		wasUsed := make(map[byte]bool)
		for in, bt := range v {
			if wasUsed[bt] {
				return fmt.Errorf("non unique byte (%v) passed to reformedOrder", in)
			}
			wasUsed[bt] = true
			eh.reformedOrder = append(eh.reformedOrder, bt)
		}
	}
	if eh.index == -1 {
		return fmt.Errorf("index out of bounds")
	}
	return nil
}

func (eh *hx256) setByte(v byte) {
	eh.raw = v
	eh.glyph = byteToString[eh.raw]
	eh.index = getIndex(eh.reformedOrder, v)

}

func (eh *hx256) AddIndex(i int) error {
	switch len(eh.reformedOrder) {
	case 0:
		c := int(eh.raw)
		return eh.Set(c + i)
	default:
		current := -1
		for c, bt := range eh.reformedOrder {
			if eh.raw != bt {
				continue
			}
			current = c
			break
		}
		return eh.Set(current + i)
	}
	return nil
}

func (eh *hx256) Add(b *hx256) error {
	for i := range b.reformedOrder {
		if eh.reformedOrder[i] != b.reformedOrder[i] {
			return fmt.Errorf("can't add: byte order must be same")
		}
	}
	sum, err := New(eh.index + b.index)
	if err != nil {
		return fmt.Errorf("can't add: %v", err.Error())
	}
	eh = sum
	return nil
}

func getIndex(reformedOrder []byte, b byte) int {
	if len(reformedOrder) == 0 {
		return int(b)
	}
	for in, bt := range reformedOrder {
		if b == bt {
			return in
		}
	}
	return -1
}

/*
0123456789ABCDEF
GHJKLMNPQRSTUVWX
YZБДЖЛЦЧШЪЫЬЭЮЯ_
!@#$%^&*()-=[]{}
|\;:'",.<>/?`~№/
abcdefghijklmnop
qrstuvwxyzбвгдёж
зиклмнптфцчшъыьэ
юя



0123456789ABCDEF
GHJKLMNPQRSTUVWX
YZabcdefghijklmn
opqrstuvwxyzΓΔΛΞ
ΠΣΦΨΩαβδεζθιλμξπ
ςστφψωБДЁЖЗИЙЛЦЧ
ШЩЪЫЬЭЮЯбвгдёжзи
йлптфцчшщъыьэюяԱ
ԲԳԴԵԷԸԹԺԻԽԾԿՀՁՂՃ
ՅՆՇՈՊՋՎՐՒՔՖԭҎҏҐґ
ҒғҔҕ⸢⸣⸤⸥⟂⟅⟆⟦⟧⟨⟩⟪
⟫⟵⟶⟷!"#$%&'()*+,
-./:;<=>?@[\]^_`
{|}~¡¢£¤¦§¨ǁǂ
"⍃","⍄","⍅","⍆","⍇","⍈","⍉","⍊","⍋","⍌","⍍","⍎","⍏","⍐","⍑","⍒",
"⌫","⌸","⌹","⌺","⌻","⌼","⌽","⌾","⍟","⍠","⍡","⍢","⍭","⍮","⍯","⍰",
"☠","☡","☢","☣","☤","☥","☦","☧","☨","☩","☪","☫","☬","☭","☮","☯",
"☰","☸","☹","☺","☻","☼","☽","☾","☿","♀","♁","♂","♃","♄","♅","♆",
"♇","♔","♕","♖","♗","♘","♙","♚","♛","♜","♝","♞","♟","♠","♡","♢",
"♣","♤","♥","♦","♧","♨","♩","♪","♫","♬","♭","♮","♯","♰","♱","♲",







*/
