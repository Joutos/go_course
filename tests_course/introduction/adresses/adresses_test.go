package adresses

import (
	"testing"
)

type TestCase struct {
	Adress   string
	Expected string
}

// TestAdressType tests the AdressType function. It checks if the function
// returns the correct address type given a string containing the address.
// The function is tested with the string "Rua 1" and the expected result is
// "Rua".
func TestAdressType(t *testing.T) {

	cases := []TestCase{
		{"Rua 1", "Rua"},
		{"Avenida 1", "Avenida"},
		{"Travessa 1", "Travessa"},
		{"Estrada 1", "Estrada"},
		{"Rodovia 1", "Rodovia"},
		{"Alameda 1", "Alameda"},
		{"Praça 1", "Praça"},
		{"Viela 1", "Viela"},
		{"Viaduto 1", "Viaduto"},
		{"Praça", "Indefinido"},
		{"", "Indefinido"},
	}

	for _, c := range cases {
		result := AdressType(c.Adress)
		if result != c.Expected {
			t.Errorf("Expected %s, got %s", c.Expected, result)
		}
	}
}
