package dice

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"sort"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// //Roll - создает и возвращает структуру из которой можно брать результат,
// //манипулировать. Нужно для одноразовых случайных бросков.
// func Roll(code string) *Dicepool {
// 	dp := Dicepool{}
// 	time.Sleep(time.Millisecondye )
// 	if dp.seed == 0 {
// 		dp.seed = time.Now().UTC().UnixNano()
// 		dp.src = rand.NewSource(dp.seed)
// 		dp.rand = *rand.New(dp.src)
// 	}
// 	dp.dice, dp.edges = decodeDiceCode(code)
// 	//rand.Seed(dp.seed)

// 	for d := 0; d < dp.dice; d++ {
// 		dp.result = append(dp.result, dp.rand.Intn(dp.edges)+1)
// 	}
// 	return &dp
// }

// //RollFromList - возвращает случайный элемент слайса на основе последовательности дайспула
// func (dp *Dicepool) RollFromList(sl []string) string {
// 	dp.result = nil
// 	dp.dice, dp.edges = 1, len(sl)
// 	dp.modPerDie = 0
// 	dp.modTotal = 0
// 	return sl[dp.rand.Intn(dp.edges)]
// }

// //RollFromList - возвращает случайный элемент слайса на основе последовательности дайспула
// func (dp *Dicepool) RollFromListInt(sl []int) int {
// 	dp.result = nil
// 	dp.dice, dp.edges = 1, len(sl)
// 	dp.modPerDie = 0
// 	dp.modTotal = 0
// 	return sl[dp.rand.Intn(dp.edges)]
// }

// func decodeDiceCode(code string) (int, int) {
// 	code = strings.ToUpper(code)
// 	data := strings.Split(code, "D")
// 	var dice int
// 	dice, _ = strconv.Atoi(data[0])
// 	if data[0] == "" {
// 		dice = 1
// 	}
// 	edges, err := strconv.Atoi(data[1])
// 	if err != nil {
// 		return 0, 0
// 	}
// 	return dice, edges
// }

// func encodeDiceCode(dice, edges int) string {
// 	return strconv.Itoa(dice) + "D" + strconv.Itoa(edges)
// }

// //////////////////////////////////////////////////////////
// //Results:

// //Result - возвращает слайс с результатами броска дайспула
// func (dp *Dicepool) Result() []int {
// 	return dp.result
// }

// //ResultIs - расшифровывает код валидных результатов и сравнивает с ними результат броска
// func (dp *Dicepool) ResultIs(code string) bool {
// 	rc := newResultCode(code)
// 	return compare(rc, dp)
// }

// func errWarn(err error) {
// 	if err != nil {
// 		fmt.Println("ERROR WARNING!!!")
// 		fmt.Println(err.Error())
// 	}
// }

// //Sum - возвращает сумму очков броска
// func (dp *Dicepool) Sum() int {
// 	sum := 0
// 	for i := 0; i < len(dp.result); i++ {
// 		sum = sum + (dp.result[i] + dp.modPerDie)
// 	}
// 	sum = sum + dp.modTotal
// 	return sum
// }

// //SumStr - возвращает сумму очков броска в виде стринга
// func (dp *Dicepool) SumStr() string {
// 	return strconv.Itoa(dp.Sum())
// }

// //ResultString - возвращает результата в виде стринга
// func (dp *Dicepool) ResultString() string {
// 	res := ""
// 	for i := 0; i < len(dp.result); i++ {
// 		res = res + strconv.Itoa(dp.result[i])
// 	}
// 	return res
// }

// //ResultTN - возвращает true если сумма броска больше/равна tn
// func (dp *Dicepool) ResultTN(tn int) bool {
// 	if dp.Sum() < tn {
// 		return false
// 	}
// 	return true
// }

// //////////////////////////////////////////////////////////
// //Actions:

// //Boon - фиксирует результат броска
// func (dp *Dicepool) Boon() *Dicepool {
// 	lowest := 0
// 	targetVal := dp.edges
// 	for i, val := range dp.result {
// 		if val < targetVal {
// 			targetVal = val
// 			lowest = i
// 		}

// 	}
// 	d1 := rand.Intn(dp.edges) + 1
// 	if d1 > targetVal {
// 		dp.result[lowest] = d1
// 	}
// 	return dp
// }

// //Bane - фиксирует результат броска
// func (dp *Dicepool) Bane() *Dicepool {
// 	highest := 0
// 	targetVal := 0
// 	for i, val := range dp.result {
// 		if val > targetVal {
// 			targetVal = val
// 			highest = i
// 		}

// 	}
// 	d1 := rand.Intn(dp.edges) + 1
// 	if d1 < targetVal {
// 		dp.result[highest] = d1
// 	}
// 	return dp
// }

// //DM - фиксирует результат броска
// func (dp *Dicepool) DM(s int) *Dicepool {
// 	dp.modTotal = s
// 	return dp
// }

// //ModPerDie - фиксирует результат броска
// func (dp *Dicepool) ModPerDie(s int) *Dicepool {
// 	dp.modPerDie = s
// 	return dp
// }

// //RerollEach - перебрасывает все unwanted
// //TODO: исключить вечную петлю
// //TODO: нужен вариант с множеством unwanted
// func (dp *Dicepool) RerollEach(unwanted int) *Dicepool {

// 	for i, val := range dp.result {
// 		for val == unwanted {
// 			val = rand.Intn(dp.edges-1) + 1
// 			dp.result[i] = val
// 		}
// 	}
// 	return dp
// }

// //ReplaceEach - заменяет все unwanted на wanted
// func (dp *Dicepool) ReplaceEach(unwanted, wanted int) *Dicepool {
// 	for i := range dp.result {
// 		for dp.result[i] == unwanted {
// 			dp.result[i] = wanted
// 		}
// 	}
// 	return dp
// }

// //ReplaceOne - меняет значение конкретного дайса
// func (dp *Dicepool) ReplaceOne(die, newVal int) *Dicepool {
// 	if len(dp.result) < die {
// 		return dp
// 	}
// 	dp.result[die] = newVal
// 	return dp
// }

// //Shout - выводит процесс в стандартный выход
// func (dp *Dicepool) Shout() *Dicepool {
// 	fmt.Println("------------------------------")
// 	fmt.Printf("Rolling: %vd%v + (%v)\n", dp.dice, dp.edges, dp.modTotal)
// 	fmt.Printf("Result: %v (%v)\n", dp.result, dp.Sum())
// 	fmt.Println("------------------------------")
// 	return dp
// }

// //ReRoll - меняет значение броска
// // func (dp *Dicepool) ReRoll() *Dicepool {
// // 	code := encodeDiceCode(dp.dice, dp.edges)
// // 	dpNew := Roll(code)
// // 	return dpNew
// // }

// //////////////////////////////////////////////////////////
// //Probe:

// //Probe - перечисляет возможные варианты
// func Probe(code string, tn int) map[int][]int {
// 	dice, edges := decodeDiceCode(code)
// 	resMap := rollCombinations(edges, dice)
// 	return resMap
// }

// //ProbeTN - оценивает вероятность достичь tn
// func ProbeTN(code string, tn int) float64 {
// 	dp := new(Dicepool)
// 	dice, edges := decodeDiceCode(code)
// 	dp.dice = dice
// 	dp.edges = edges
// 	var positiveOutcome int
// 	for i := 0; i < dice; i++ {
// 		dp.result = append(dp.result, 1)
// 	}
// 	totalRange := int(math.Pow(float64(dp.edges), float64(dp.dice)))
// 	resMap := rollCombinations(dp.edges, dp.dice)
// 	for _, v := range resMap {
// 		tdp := Dicepool{}
// 		tdp.result = v
// 		if tdp.ResultTN(tn) {
// 			positiveOutcome++
// 		}
// 	}
// 	res := float64(positiveOutcome) / float64(totalRange)
// 	return res
// }

// func rollCombinations(max, len int) map[int][]int {
// 	resMap := make(map[int][]int)
// 	var sl []int
// 	for i := 0; i < len; i++ {
// 		sl = append(sl, 1)
// 	}
// 	totalRange := int(math.Pow(float64(max), float64(len)))
// 	resMap[0] = sl
// 	for i := 0; i < totalRange; i++ {
// 		activeDie := 0
// 		for activeDie < len {
// 			if sl[activeDie] < max {
// 				sl[activeDie]++
// 				for j := range sl {
// 					resMap[i+1] = append(resMap[i+1], sl[j])
// 				}
// 				break
// 			} else {
// 				sl[activeDie] = 1
// 				activeDie++
// 			}
// 			if activeDie > len {
// 				break
// 			}
// 		}
// 	}
// 	return resMap
// }

// //////////////////////////////////////////////////////////
// //QuickRolls:

// //RollD66 - Возвращает результат 2d6 в виде string
// func RollD66() string {
// 	return Roll("2d6").ResultString()

// }

// //Roll1D -
// func Roll1D(dm ...int) int {
// 	mod := 0
// 	if len(dm) > 0 {
// 		mod = dm[0]
// 	}
// 	return Roll("1d6").DM(mod).Sum()
// }

// //RollD3 -
// func RollD3(dm ...int) int {
// 	mod := 0
// 	if len(dm) > 0 {
// 		mod = dm[0]
// 	}
// 	return Roll("1d3").DM(mod).Sum()
// }

// //Roll2D -
// func Roll2D(dm ...int) int {
// 	mod := 0
// 	if len(dm) > 0 {
// 		mod = dm[0]
// 	}
// 	return Roll("2d6").DM(mod).Sum()
// }

// //Roll3D -
// func Roll3D(dm ...int) int {
// 	mod := 0
// 	if len(dm) > 0 {
// 		mod = dm[0]
// 	}
// 	return Roll("3d6").DM(mod).Sum()
// }

// //Roll4D -
// func Roll4D(dm ...int) int {
// 	mod := 0
// 	if len(dm) > 0 {
// 		mod = dm[0]
// 	}
// 	return Roll("4d6").DM(mod).Sum()
// }

// //Roll5D -
// func Roll5D(dm ...int) int {
// 	mod := 0
// 	if len(dm) > 0 {
// 		mod = dm[0]
// 	}
// 	return Roll("5d6").DM(mod).Sum()
// }

// func Flux() int {
// 	d1 := Roll1D()
// 	d2 := Roll1D()
// 	return d1 - d2
// }

// func FluxMicro() int {
// 	d1 := Roll1D()
// 	switch d1 {
// 	default:
// 		return 0
// 	case 1:
// 		return -1
// 	case 6:
// 		return 1
// 	}
// }

// func FluxGOOD() int {
// 	d1 := Roll1D()
// 	d2 := Roll1D()
// 	if d1 >= d2 {
// 		return d1 - d2
// 	}
// 	return d2 - d1
// }

// func FluxBAD() int {
// 	d1 := Roll1D()
// 	d2 := Roll1D()
// 	if d1 <= d2 {
// 		return d1 - d2
// 	}
// 	return d2 - d1
// }

// ////////////////
// //ResultCodeReader

// type resultCode struct {
// 	codeBody      string
// 	codeType      string //Asend/Desend/slice
// 	valid         bool
// 	compareValues []int
// 	err           error
// }

// func compare(rc resultCode, dp *Dicepool) bool {
// 	if !rc.valid {
// 		return false
// 	}
// 	switch rc.codeType {
// 	case "+":
// 		if dp.Sum() >= rc.compareValues[0] {
// 			return true
// 		}
// 	case "-":
// 		if dp.Sum() <= rc.compareValues[0] {
// 			return true
// 		}
// 	default:
// 		for i := rc.compareValues[0]; i <= rc.compareValues[1]; i++ {
// 			if i == dp.Sum() {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func newResultCode(code string) resultCode {
// 	rc := resultCode{}
// 	rc.codeBody = code
// 	last := string(rc.codeBody[len(rc.codeBody)-1:])
// 	switch last {
// 	case "+", "-":
// 		rc.codeType = last
// 	default:
// 		last = ""
// 		rc.codeType = "slice"
// 	}
// 	rc.codeBody = strings.TrimSuffix(rc.codeBody, last)
// 	codeParts := strings.Split(rc.codeBody, " ")
// 	for i, part := range codeParts {
// 		pt, err := strconv.Atoi(part)
// 		switch {
// 		case err == nil:
// 			rc.compareValues = append(rc.compareValues, pt)
// 		default:
// 			rc.err = errors.New("can't read resultCode value from part " + strconv.Itoa(i) + ": " + err.Error())
// 			return rc
// 		}
// 	}
// 	if len(rc.compareValues) != 1 && (rc.codeType == "+" || rc.codeType == "-") {
// 		rc.err = errors.New("code parsing incorect")
// 		return rc
// 	}
// 	if len(rc.compareValues) == 1 {
// 		rc.compareValues = append(rc.compareValues, rc.compareValues[0])
// 	}
// 	sort.Sort(sort.IntSlice(rc.compareValues))
// 	rc.valid = true
// 	return rc
// }

// /*
// s01e03: Наводя мосты.
// За получив на борт 761 беженцев, команда Откровения оказалась в достаточно щепетильной ситуации. В большинстве своём Аликай расположены к неизвестным и странным пришельцам позитивно: еще бы они спасли Бридеров!
// Однако не все жуки разделают этот оптимизм. Как позже, благодаря уже освоевшемуся среди людей Акранике, выяснилось что Аликай имеют очень жесткую и тесную социальную структуру общества. И те из спасенных кто лишился своих Бридеров Рискуют стать Отверженными - худшее из возможных исходов для Аликай.
// С дальностью прыжка откровения в 4 парсека. Понадобится всего 1 неделя, чтобы добраться до скопления звезд, в котором находится цивилизация Аликай. Что там ждет Откровение? Как их встретит раса до сих пор ни разу не видевшая инопланетян?
// ......
// -Капитан! Все проверки пройдены, корабль готов ко входу в гипер пространство.
// -Хорошо. Всему экипажу. Приготовиться. Мы будем прыгать!

// */

/*
examples:
1d6			die=1 edges=6 dm=0
2d6+7		die=2 edges=6 dm=7
2d6-7		die=2 edges=6 dm=-7
6DD
2d+3
Roll("2D+8")

*/
