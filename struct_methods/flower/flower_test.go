package flower

import "testing"

func TestGetName(t *testing.T) {
	f := &Flower{Name: "Rose"}
	if f.GetName() != "Rose" {
		t.Errorf("expected Rose, got %s", f.GetName())
	}
}
