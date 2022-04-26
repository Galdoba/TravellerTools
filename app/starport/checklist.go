package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func Checklist(c *cli.Context) error {
	return fmt.Errorf("Not Implemented")
}

/*
World  : [NAME] ([SECTOR] [HEX])
UWP    : XXXXXXX-X
TC/Rem : [Trade codes and remarks] ([ZONE])
-------------------------------------------
запросы:
Размер
Защищенность
Посадочные площадки (свободные/занятые сейчас)
верфи/топливо/склады

if TL8+ {
Computer network Available.
Updating ship registry logs...
You have [3D-14] new message(s)...  -лайф эвент контакта или автогенерированные локальные квесты предложения
}
Landing...



Является ли мир часть торгового пути?
Источник - товары из этого мира уходят в космос
Рынок - товары в этот мир приходят из источника
Транзитный пункт - корабли просо проходят мимо
Захолустье - нет причин лететь в этот мир



*/
