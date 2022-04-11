package balancingact

type faction struct {
	origin      string
	factionType int //Political, Economic
	name        string
	balance     float64
	control     map[*planet]*asset
	leader      *pawn
	agents      map[int]*pawn
}
