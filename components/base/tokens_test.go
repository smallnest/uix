package tokens

import "testing"

func TestLightAndDark(t *testing.T) {
	light, dark := Light(), Dark()
	if light.Dark {
		t.Fatalf("light theme must not be dark")
	}
	if !dark.Dark {
		t.Fatalf("dark theme must be dark")
	}
	if light.Background == dark.Background {
		t.Fatalf("palettes must differ, both are %v", light.Background)
	}
	if dark.Radius != 8 {
		t.Fatalf("palette radius not applied: got %v", dark.Radius)
	}
}

func TestFor(t *testing.T) {
	if !For(true).Dark {
		t.Fatalf("For(true) must return the dark theme")
	}
	if For(false).Dark {
		t.Fatalf("For(false) must return the light theme")
	}
}
