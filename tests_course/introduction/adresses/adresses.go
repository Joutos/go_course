package adresses

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AdressType receives a string containing an address and returns the type of address.
// The types supported are: Rua, Avenida, Travessa, Estrada, Rodovia, Alameda, Praça, Viela, Viaduto.
// If the address type is not found, the function returns "Indefinido".
func AdressType(adress string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	adress = caser.String(adress)

	types := []string{
		"Rua",
		"Avenida",
		"Travessa",
		"Estrada",
		"Rodovia",
		"Alameda",
		"Praça",
		"Viela",
		"Viaduto",
	}

	rd := strings.Split(adress, " ")[0]

	for _, t := range types {
		if t == rd {
			return rd
		}
	}
	return "Indefinido"

}
