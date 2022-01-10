package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//Dicepool -
type Dicepool struct {
	dice       int
	edges      int
	modPerDie  int
	modTotal   int
	seed       int64
	result     []int
	randomSeed bool
	boon       bool
	bane       bool
	src        rand.Source
	rand       rand.Rand
}

//New - Creates Dicepool object with seed from current time
func New() *Dicepool {
	dp := Dicepool{}
	dp.Reset()
	return &dp
}

//SetSeed - фиксирует результат броска
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

//seedFromString - задает Seed по ключу key
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

//Roll - Делает бросок
func (dp *Dicepool) Roll(code string) *Dicepool {
	dp.result = nil
	dp.dice, dp.edges = decodeDiceCode(code)
	dp.modPerDie = 0
	dp.modTotal = 0
	for d := 0; d < dp.dice; d++ {
		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
	}
	return dp
}

//Flux - Дает Flux
func (dp *Dicepool) Flux() int {
	d1 := dp.Roll("1d6").Sum()
	d2 := dp.Roll("1d6").Sum()
	return d1 - d2
}

//Sum - возвращает сумму очков броска
func (dp *Dicepool) Sum() int {
	sum := 0
	for i := 0; i < len(dp.result); i++ {
		sum = sum + (dp.result[i] + dp.modPerDie)
	}
	sum = sum + dp.modTotal
	return sum
}

//SumStr - возвращает сумму очков броска в виде стринга
func (dp *Dicepool) SumStr() string {
	return strconv.Itoa(dp.Sum())
}

func decodeDiceCode(code string) (int, int) {
	code = strings.ToUpper(code)
	data := strings.Split(code, "D")
	var dice int
	dice, _ = strconv.Atoi(data[0])
	if data[0] == "" {
		dice = 1
	}
	edges, err := strconv.Atoi(data[1])
	if err != nil {
		switch {
		case strings.Contains(data[1], "+"):
			fmt.Println("-------")
		}
		return 0, 0
	}
	return dice, edges
}

func (d *Dicepool) TreatAasB(a, b int) {
	for i, r := range d.result {
		if r == a {
			d.result[i] = b
		}
	}
}

/*
чего ку
амиксин
ринза
*/
