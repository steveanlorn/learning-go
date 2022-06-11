package count

// Counter ...
type Counter struct {
	count int
}

// NewCounter ...
func NewCounter() *Counter {
	return new(Counter)
}

// Up ...
func (c *Counter) Up() int {
	c.count++
	return c.count
}

// Down ...
func (c *Counter) Down() int {
	c.count--
	return c.count
}

func (c *Counter) GetCount() int {
	return c.count
}
