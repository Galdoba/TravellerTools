package balancingact

const (
	ZeroValue = iota
	Leader
	Agent
	WIL
	INT
	EDU
	CHR
	Administration
	CovertOps
	Economics
	Politics
	Military
	Solidarity
	Wealth
	Expansion
	Might
	Development
	Starport
	Population
	Govr
	Law
	TimeFactor_Day
	TimeFactor_Week
	TimeFactor_Month
	TimeFactor_Year
	STPRT_A
	STPRT_B
	STPRT_C
	STPRT_D
	STPRT_E
	STPRT_X
	TASK_Changing_the_Law_Level_of_a_World
	TASK_Changing_the_Might__of_a_world_which_you_do_not_rule
	UnlistedValue
)

//фактически это игрок
type ControlGroup struct {
	name             string // название (служет как уникальный ID)
	organisationType int    // тип фракции (определяет доступные действия ее участников)
	asset            asset  // владения (влияние, персонал, армия, флот и тд.)
}

// владения (влияние, персонал, армия, флот и тд.)
type asset struct {
	personal  map[int]*pawn
	influence map[*planet]share
	//fleet     map[int]int //TODO: отдельная структура
	//army      map[int]int //TODO: отдельная структура
}

type share struct {
	polity        *planet
	influenceType int
	fracture      float64
	profit        float64
	expenses      float64
}
