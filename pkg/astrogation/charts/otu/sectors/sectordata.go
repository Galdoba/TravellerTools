package sectors

import (
	"fmt"
	"strings"
)

func NameAbb(x, y int) (string, string) {
	name := "Uncharted " + fmt.Sprintf("%v%v", x, y)
	for _, sect := range sectors() {
		if x == sect.x && y == sect.y {
			name = sect.fullName
		}
	}
	return name, abb(name)
}

type sector struct {
	x        int
	y        int
	fullName string
}

func abb(name string) string {
	name = strings.ReplaceAll(name, "'", " ")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, " ", "")
	name = name[:4]
	return name
}

func sectors() []sector {
	var s []sector
	s = append(s, sector{-4, -1, "Spinward Marches"})
	s = append(s, sector{-3, -1, "Deneb"})
	s = append(s, sector{-2, -1, "Corridor"})
	s = append(s, sector{-1, -1, "Vland"})
	s = append(s, sector{0, -1, "Lishun"})
	s = append(s, sector{1, -1, "Antares"})
	s = append(s, sector{2, -1, "Empty Quarter"})

	s = append(s, sector{-4, 0, "Trojan Reach"})
	s = append(s, sector{-3, 0, "Reft Sector"})
	s = append(s, sector{-2, 0, "Gushemege"})
	s = append(s, sector{-1, 0, "Dagudashag"})
	s = append(s, sector{0, 0, "Core"})
	s = append(s, sector{1, 0, "Fornast"})
	s = append(s, sector{2, 0, "Ley Sector"})

	s = append(s, sector{-4, 1, "Riftspan Reaches"})
	s = append(s, sector{-3, 1, "Verge"})
	s = append(s, sector{-2, 1, "Ilelish"})
	s = append(s, sector{-1, 1, "Zarushagar"})
	s = append(s, sector{0, 1, "Massila"})
	s = append(s, sector{1, 1, "Delpi"})
	s = append(s, sector{2, 1, "Glimmerdrift Reaches"})

	s = append(s, sector{-4, 2, "Hlakhoi"})
	s = append(s, sector{-3, 2, "Ealiyasiyw"})
	s = append(s, sector{-2, 2, "Reaver's Deep"})
	s = append(s, sector{-1, 2, "Daibei"})
	s = append(s, sector{0, 2, "Diaspora"})
	s = append(s, sector{1, 2, "Old Expanses"})
	s = append(s, sector{2, 2, "Hinter-Worlds"})

	s = append(s, sector{-4, 3, "Staihaia'yo"})
	s = append(s, sector{-3, 3, "Iwahfuah"})
	s = append(s, sector{-2, 3, "Dark Nebula"})
	s = append(s, sector{-1, 3, "Magyar"})
	s = append(s, sector{0, 3, "Solomani Rim"})
	s = append(s, sector{1, 3, "Alpha Crucis"})
	s = append(s, sector{2, 3, "Spica"})
	return s
}
