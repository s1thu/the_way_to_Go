package car

import "testing"

func TestCar_SetName(t *testing.T) {
	c := &Car{}
	c.SetName("Honda")
	if c.Name != "Honda" {
		t.Errorf("expected Honda, got %s", c.Name)
	}
}

func TestCar_SetYear(t *testing.T) {
	c := &Car{}
	c.SetYear(2022)
	if c.Year != 2022 {
		t.Errorf("expected 2022, got %d", c.Year)
	}
}
