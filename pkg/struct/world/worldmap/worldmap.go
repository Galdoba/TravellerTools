package worldmap

type worldmap struct {
	Triangle map[int]WorldTriangle
}

type WorldTriangle struct {
	ID    int
	ANeib int
	BNeib int
	CNeib int
}

func d20tod66(i int) int {
	switch i {
	case 1:
		return 11
	case 2:
		return 12
	case 3:
		return 13
	case 4:
		return 14
	case 5:
		return 15
	case 6:
		return 21
	case 7:
		return 22
	case 8:
		return 23
	case 9:
		return 24
	case 10:
		return 25
	case 11:
		return 31
	case 12:
		return 32
	case 13:
		return 33
	case 14:
		return 34
	case 15:
		return 35
	case 16:
		return 41
	case 17:
		return 42
	case 18:
		return 43
	case 19:
		return 44
	case 20:
		return 45
	}
	return -1
}

func NewTriangle(ID int) WorldTriangle {
	tr := WorldTriangle{}
	tr.ID = ID
	switch ID {
	case 11:
		tr.ANeib = 21
		tr.BNeib = 15
		tr.CNeib = 12
	case 12:
		tr.ANeib = 22
		tr.BNeib = 11
		tr.CNeib = 13
	case 13:
		tr.ANeib = 23
		tr.BNeib = 12
		tr.CNeib = 14
	case 14:
		tr.ANeib = 24
		tr.BNeib = 13
		tr.CNeib = 15
	case 15:
		tr.ANeib = 23
		tr.BNeib = 12
		tr.CNeib = 14
	case 21:
		tr.ANeib = 11
		tr.BNeib = 31
		tr.CNeib = 35
	case 22:
		tr.ANeib = 12
		tr.BNeib = 32
		tr.CNeib = 31
	case 23:
		tr.ANeib = 13
		tr.BNeib = 33
		tr.CNeib = 32
	case 24:
		tr.ANeib = 14
		tr.BNeib = 34
		tr.CNeib = 33
	case 25:
		tr.ANeib = 15
		tr.BNeib = 35
		tr.CNeib = 34
	case 31:
		tr.ANeib = 41
		tr.BNeib = 22
		tr.CNeib = 21
	case 32:
		tr.ANeib = 42
		tr.BNeib = 23
		tr.CNeib = 22
	case 33:
		tr.ANeib = 43
		tr.BNeib = 24
		tr.CNeib = 23
	case 34:
		tr.ANeib = 44
		tr.BNeib = 25
		tr.CNeib = 24
	case 35:
		tr.ANeib = 45
		tr.BNeib = 21
		tr.CNeib = 25
	case 41:
		tr.ANeib = 31
		tr.BNeib = 42
		tr.CNeib = 45
	case 42:
		tr.ANeib = 32
		tr.BNeib = 43
		tr.CNeib = 41
	case 43:
		tr.ANeib = 33
		tr.BNeib = 44
		tr.CNeib = 42
	case 44:
		tr.ANeib = 34
		tr.BNeib = 45
		tr.CNeib = 43
	case 45:
		tr.ANeib = 35
		tr.BNeib = 41
		tr.CNeib = 44
	}
	return tr
}
