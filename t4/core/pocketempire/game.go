package pocketempire

import (
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/empire"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/empire/worldcharacter"
)

/*FLOW


ежедневно  :
	-просчитать запланированные события
еженедельно:
	-переместить юниты
	-просчитать бой
ежемесячно :
	-оплатить содержание юнитов
ежегодно   :
	-пересчитать бюджет
	-культурный дрифт
	-рост населения
	-запустить планетарный проект
	-наметить мета-задачи
	-наметить случайные события



*/

type Game struct {
	CurrentDate calendar.Date
	Universe    map[hexagon.Hexagon]worldcharacter.World
	Empires     map[int]*empire.PocketEmpire
}

type eventPlanner struct {
	current       calendar.Date
	plannedEvents map[uint64]event
}

type event struct {
	descr string
}
