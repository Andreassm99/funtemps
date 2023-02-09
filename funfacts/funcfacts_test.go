package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string   // her må du skrive riktig type for input
		want  []string // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	//tests := []test{
	//  {input: , want: },
	//}

	tests := []test{
		{input: "Sun", want: []string{"Temperaturen i solens kjerne er 15 000 000C*", "Temperaturen på ytre laget av Solen er 5 500C*"}},
		{input: "Luna", want: []string{"Temperaturen på månens overflate om natten er -183C*", "Temperaturen på månens overflate om dagen er 379C*"}},
		{input: "Terra", want: []string{"Høyeste temperaturen på jordens overflate er 56,7C*", "Laveste temperatur målt på jordens overlate er -89,9C*", "Temperatur i jordens indre kjerne er 9 108C*"}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
