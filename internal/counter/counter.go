package counter

type counter struct {
	value int
}

func New() *counter {
	return &counter{0}
}

type Counter interface {
	Set(int)
	Increase()
	Decrease()
	Add(int)
	Current() int
}

func (c *counter) Set(i int) {
	c.value = i
}

func (c *counter) Increase() {
	c.value++
}

func (c *counter) Add(v int) {
	c.value += v
}

func (c *counter) Decrease() {
	c.value--
}

func (c *counter) Current() int {
	return c.value
}

func Delete(c Counter) {
	c = nil
}
