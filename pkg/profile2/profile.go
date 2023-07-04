package profile2

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

type uniProfile struct {
	data          map[string]ehex.Ehex
	profileType   string
	exclusiveKeys []string
}

type Profile interface {
	Create(string) error
	Read(string) (ehex.Ehex, error)
	Update(string, ehex.Ehex) error
	Delete(string) error
	Keys() []string
	Type() string
}

func New() *uniProfile {
	pr := uniProfile{}
	pr.data = make(map[string]ehex.Ehex)
	pr.profileType = "simple"
	return &pr
}

//Create - создает запись в профайле по ключу key
//возвращает ошибку, если ключ не подходит или занят
func (pr *uniProfile) Create(key string) error {
	if len(pr.exclusiveKeys) != 0 {
		if !inSlice(pr.exclusiveKeys, key) {
			return fmt.Errorf("cann't create entry: key '%v' is not allowed", key)
		}
	}
	if _, ok := pr.data[key]; ok {
		return fmt.Errorf("cann't create entry: key '%v' already created", key)
	}
	pr.data[key] = ehex.New()
	return nil
}

//Read - возвращает значение по ключу key
func (pr *uniProfile) Read(key string) (ehex.Ehex, error) {
	if len(pr.exclusiveKeys) != 0 {
		if !inSlice(pr.exclusiveKeys, key) {
			return nil, fmt.Errorf("cann't read entry: key '%v' is not allowed", key)
		}
	}
	if _, ok := pr.data[key]; !ok {
		return nil, fmt.Errorf("cann't read entry: key '%v' is not created", key)
	}
	return pr.data[key], nil
}

//Update - заменяет имеющееся значение новым
func (pr *uniProfile) Update(key string, val ehex.Ehex) error {
	if _, ok := pr.data[key]; !ok {
		return fmt.Errorf("cann't update entry: key '%v' is not created", key)
	}
	pr.data[key] = val
	return nil
}

//Delete - удаляет имеющееся занчение с ключем key
func (pr *uniProfile) Delete(key string) error {
	if len(pr.exclusiveKeys) != 0 {
		if !inSlice(pr.exclusiveKeys, key) {
			return fmt.Errorf("cann't delete entry: key '%v' is not allowed", key)
		}
	}
	if _, ok := pr.data[key]; !ok {
		return fmt.Errorf("cann't delete entry: key '%v' is not created", key)
	}
	delete(pr.data, key)
	return nil
}

//Keys - возвращает разрешенные ключи
//если список пуст - разрешены все
func (pr *uniProfile) Keys() []string {
	return pr.exclusiveKeys
}

//Type - возвращает тип профайла
func (pr *uniProfile) Type() string {
	return pr.profileType
}

//control//////////////

func (pr *uniProfile) Narrow(ptype string, keys []string) *uniProfile {
	pr.exclusiveKeys = keys
	pr.profileType = ptype
	return pr
}

func Append(prf1, prf2 Profile) Profile {
	newPrf := New()
	for _, k := range prf1.Keys() {
		val, _ := prf1.Read(k)
		if newPrf.Create(k) == nil {
			newPrf.Update(k, val)
		}
	}
	for _, k := range prf2.Keys() {
		val, _ := prf2.Read(k)
		if newPrf.Create(k) == nil {
			newPrf.Update(k, val)
		}
	}
	return newPrf
}

//helpers//////////////////////

func matchByKeys(bigger, smaller Profile) bool {
	if len(bigger.Keys()) == 0 {
		return true
	}
	for _, keyS := range smaller.Keys() {
		if !inSlice(bigger.Keys(), keyS) {
			return false
		}
	}
	return true
}

func inSlice(sl []string, s string) bool {
	for _, ss := range sl {
		if s == ss {
			return true
		}
	}
	return false
}
