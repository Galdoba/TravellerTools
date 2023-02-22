package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Dicepool -
type Dicepool struct {
	dice        int
	edges       int
	modPerDie   int
	modTotal    int
	seed        int64
	result      []int
	randomSeed  bool
	boon        bool
	bane        bool
	destructive bool
	src         rand.Source
	rand        rand.Rand
	vocal       bool
	err         error
}

// New - Creates Dicepool object with seed from current time
func New() *Dicepool {
	dp := Dicepool{}
	dp.Reset()
	return &dp
}

// SetSeed - фиксирует результат броска
func (dp *Dicepool) SetSeed(key interface{}) *Dicepool {
	var s int64
	switch key.(type) {
	default:
		return dp
	case int:
		s = int64(key.(int))
	case int32:
		s = int64(key.(int32))
	case int64:
		s = key.(int64)
	case string:
		s = seedFromString(key.(string))
		if key.(string) == "" {
			s = time.Now().UnixNano()
		}
	}
	dp.seed = s
	dp.randomSeed = false
	dp.result = nil
	dp.src = rand.NewSource(dp.seed)
	dp.rand = *rand.New(dp.src)
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
	}
	return dp
}

func (dp *Dicepool) Reset() {
	dp = reset(dp)
}

func reset(dp *Dicepool) *Dicepool {
	dp.randomSeed = true
	dp.seed = time.Now().UTC().UnixNano()
	dp.src = rand.NewSource(dp.seed)
	dp.rand = *rand.New(dp.src)
	return dp
}

// seedFromString - задает Seed по ключу key
func seedFromString(key string) int64 {
	bytes := []byte(key)
	var seed int64
	for i := range bytes {
		r := rune(bytes[i])
		p := int64(r) * int64(i+1)
		seed = seed + p
		// if i > 255 { Возможно понадобится ограничитель
		// 	break
		// }
	}
	return seed
}

// Roll - Делает бросок
func (dp *Dicepool) Roll(code string) *Dicepool {
	dp.result = nil
	dp.modPerDie = 0
	dp.modTotal = 0
	dp.destructive = false
	if dp.err = dp.decodeDiceCode(code); dp.err != nil {
		return dp
	}
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
	}
	if dp.vocal {
		fmt.Printf("   Rolled [%v + %v]: result is %v\n", code, dp.modTotal, dp.result)
		fmt.Printf("   Code [%v]: %v\n", code, dp.edges)
	}
	return dp
}

func (dp *Dicepool) Sroll(code string) int {
	return dp.Roll(code).Sum()
}

func (dp *Dicepool) Result() []int {
	return dp.result
}

func (dp *Dicepool) Vocal() {
	if dp.vocal == true {
		dp.vocal = false
		return
	}
	if dp.vocal == false {
		dp.vocal = true
		return
	}
}

// DrawData - возвращает количество дайсов и их грани
func (dp *Dicepool) DrawData() (int, int) {
	return dp.dice, dp.edges
}

// ModTotal - возвращает общий модификатор
func (dp *Dicepool) ModTotal() int {
	return dp.modTotal
}

// DM - добавляет число к общей сумме результата
func (dp *Dicepool) DM(dm int) *Dicepool {
	dp.modTotal = dp.modTotal + dm
	return dp
}

// Flux - Дает Flux
func (dp *Dicepool) Flux() int {
	d1 := dp.Roll("1d6").Sum()
	d2 := dp.Roll("1d6").Sum()
	return d1 - d2
}

// Sum - возвращает сумму очков броска
func (dp *Dicepool) Sum() int {
	sum := 0
	for i := 0; i < len(dp.result); i++ {
		sum = sum + (dp.result[i] + dp.modPerDie)
	}
	sum = sum + dp.modTotal
	if dp.destructive {
		sum = sum * 10
	}
	return sum
}

// SumStr - возвращает сумму очков броска в виде стринга
func (dp *Dicepool) SumStr() string {
	return strconv.Itoa(dp.Sum())
}

func (dp *Dicepool) decodeDiceCode(code string) error {
	code = strings.ToUpper(code)
	dest := strings.TrimSuffix(code, "DD")
	switch {
	case code != dest:
		return dp.decodeDestructive(dest)
	case code == dest:
		return dp.decodeNormal(code)
	}
	return fmt.Errorf("decoding failed due to unknown reason")
}

func (dp *Dicepool) decodeDestructive(code string) error {
	diceNum, err := strconv.Atoi(code)
	if err != nil {
		return fmt.Errorf("cannot parse dice code '%v'", code)
	}
	dp.dice = diceNum
	dp.edges = 6
	dp.destructive = true
	return nil
}

func (dp *Dicepool) decodeNormal(code string) error {
	p1 := strings.Split(code, "D")
	if len(p1) != 2 {
		return fmt.Errorf("cannot parse dice code '%v'", code)
	}
	d, err := strconv.Atoi(p1[0])
	if err != nil {
		return fmt.Errorf("cannot parse dice code '%v'", code)
	}
	dp.dice = d
	switch {
	///////////////////no mod
	case !strings.Contains(p1[1], "+") && !strings.Contains(p1[1], "-"):
		edges, err := strconv.Atoi(p1[1])
		if err != nil {
			return err
		}
		dp.edges = edges
		///////////////////positive mod
	case strings.Contains(p1[1], "+"):
		p2 := strings.Split(p1[1], "+")
		for i, v := range p2 {
			n, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("cannot parse dice code '%v'", code)
			}
			switch i {
			case 0:
				dp.edges = n
			case 1:
				dp.modTotal = dp.modTotal + n
			}
		}
		///////////////////negative mod
	case strings.Contains(p1[1], "-"):
		p2 := strings.Split(p1[1], "-")
		for i, v := range p2 {
			n, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("cannot parse dice code '%v'", code)
			}
			switch i {
			case 0:
				dp.edges = n
			case 1:
				dp.modTotal = dp.modTotal - n
			}
		}
	}
	return nil
}

func (d *Dicepool) TreatAasB(a, b int) {
	for i, r := range d.result {
		if r == a {
			d.result[i] = b
		}
	}
}

// //////////////////
// Pick - возвращает случайный элемент из слайса и его номер
func (d *Dicepool) Pick(slice []interface{}) (int, interface{}) {
	l := fmt.Sprintf("%v", len(slice))
	r := d.Roll("1d" + l).DM(-1).Sum()
	return r, slice[r]
}

func (d *Dicepool) PickStr(slice []string) (int, string) {
	l := fmt.Sprintf("%v", len(slice))
	r := d.Roll("1d" + l).DM(-1).Sum()
	return r, slice[r]
}

func (d *Dicepool) PickStrOnly(slice []string) string {
	l := fmt.Sprintf("%v", len(slice))
	r := d.Roll("1d" + l).DM(-1).Sum()
	return slice[r]
}

/*
чего ку
амиксин
ринза
*/
