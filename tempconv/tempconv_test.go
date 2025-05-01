package tempconv

import (
	"testing"
)

// TestCToF testa a conversão de Celsius para Fahrenheit
func TestCToF(t *testing.T) {
	tests := []struct {
		celsius  Celsius
		wantFahr Fahrenheit
	}{
		{Celsius(0), Fahrenheit(32)},
		{Celsius(100), Fahrenheit(212)},
		{Celsius(-40), Fahrenheit(-40)},
		{Celsius(37), Fahrenheit(98.6)},
	}

	for _, tt := range tests {
		got := CToF(tt.celsius)
		if diff := got - tt.wantFahr; diff < -0.01 || diff > 0.01 {
			t.Errorf("CToF(%v) = %v, want %v", tt.celsius, got, tt.wantFahr)
		}
	}
}

// TestFToC testa a conversão de Fahrenheit para Celsius
func TestFToC(t *testing.T) {
	tests := []struct {
		fahrenheit  Fahrenheit
		wantCelsius Celsius
	}{
		{Fahrenheit(32), Celsius(0)},
		{Fahrenheit(212), Celsius(100)},
		{Fahrenheit(-40), Celsius(-40)},
		{Fahrenheit(98.6), Celsius(37)},
	}

	for _, tt := range tests {
		got := FToC(tt.fahrenheit)
		if diff := got - tt.wantCelsius; diff < -0.01 || diff > 0.01 {
			t.Errorf("FToC(%v) = %v, want %v", tt.fahrenheit, got, tt.wantCelsius)
		}
	}
}

// TestStringMethods testa o método String das temperaturas
func TestStringMethods(t *testing.T) {
	c := Celsius(25)
	if c.String() != "25°C" {
		t.Errorf("Celsius.String() = %v, want %v", c.String(), "25°C")
	}

	f := Fahrenheit(77)
	if f.String() != "77°F" {
		t.Errorf("Fahrenheit.String() = %v, want %v", f.String(), "77°F")
	}
}

// TestConstants testa se as constantes estão com os valores corretos
func TestConstants(t *testing.T) {
	if AbsoluteZeroC != -273.15 {
		t.Errorf("AbsoluteZeroC = %v, want -273.15", AbsoluteZeroC)
	}
	if FreezingC != 0 {
		t.Errorf("FreezingC = %v, want 0", FreezingC)
	}
	if BoilingC != 100 {
		t.Errorf("BoilingC = %v, want 100", BoilingC)
	}
}
