package car

type Car struct {
	Name string
	Year int
}

func (c *Car) SetName(name string) {
	c.Name = name
}

func (c *Car) SetYear(year int) {
	c.Year = year
}
