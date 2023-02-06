package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	//return 56.67
	return (value - 32) * 5 / 9
}

// De andre konverteringsfunksjonene implementere her
// ...

func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9.0 / 5.0) + 32
}

func KelvinToFahrenheit(value float64) float64 {
	return (value * 9.0 / 5.0) - 459.67
}

func FahrenheitToKelvin(value float64) float64 {
	return (value + 459.67) * 5.0 / 9.0
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}
