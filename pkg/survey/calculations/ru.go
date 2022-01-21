package calculations

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/ehex"
)

func RU(econ string) int {

	hex := strings.Split(econ, "")
	r := ehex.New().Set(hex[1])
	l := ehex.New().Set(hex[2])
	i := ehex.New().Set(hex[3])
	e, _ := strconv.Atoi(hex[4] + hex[5])
	return r.Value() * l.Value() * i.Value() * e
}
