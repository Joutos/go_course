package shapes

import "testing"

func TestArea(t *testing.T) {
	t.Run("Área do Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 15}
		expected := 150.0
		area := ret.Area()

		if area != expected {
			t.Errorf("Expected %0.2f, got %0.2f", expected, area)
		}
	})

	t.Run("Área do Círculo", func(t *testing.T) {
		circ := Circulo{10}
		expected := 314.1592653589793
		area := circ.Area()

		if area != expected {
			t.Errorf("Expected %0.2f, got %0.2f", expected, area)
		}
	})

	t.Run("Área do Triângulo", func(t *testing.T) {
		tri := Triangulo{12, 6}
		expected := 36.0
		area := tri.Area()

		if area != expected {
			t.Errorf("Expected %0.2f, got %0.2f", expected, area)
		}
	})
}
