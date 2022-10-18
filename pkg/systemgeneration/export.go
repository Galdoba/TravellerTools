package systemgeneration

func Export(class string, num int, size string) *star {
	st := star{
		class:          class,
		num:            num,
		size:           size,
		orbit:          map[float64]StellarBody{},
		orbitDistances: []float64{},
	}
	st.LoadValues()
	return &st
}

func (st *star) Class() string {
	return st.class
}

func (st *star) InnerLimit() float64 {
	return st.innerLimit
}
func (st *star) HabitabilityLow() float64 {
	return st.habitableLow
}
func (st *star) HabitabilityHi() float64 {
	return st.habitableHigh
}
func (st *star) SnowLine() float64 {
	return st.snowLine
}
func (st *star) OuterLimit() float64 {
	return st.outerLimit
}
func (st *star) Mass() float64 {
	return st.mass
}
func (st *star) Luminocity() float64 {
	return st.luminocity
}
func (st *star) Position() string {
	return st.distanceType
}
