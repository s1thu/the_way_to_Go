package flower

type Flower struct {
	Name  string
	Color string
}

func (f *Flower) SetName(name string) {
	f.Name = name
}
func (f *Flower) SetColor(color string) {
	f.Color = color
}
func (f *Flower) GetName() string {
	return f.Name
}

func (f *Flower) GetColor() string {
	return f.Color
}
