package conv

import "math"

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
  FtoC := (value - 32) * (5.0 / 9.0)
	return (math.Round(FtoC * 100) / 100)
}

// De andre konverteringsfunksjonene implementere her
// ...

func FahrenheitToKelvin(value float64) float64 {
  FtoK := (value - 32) * (5.0 / 9.0) + 273.15
  return (math.Round(FtoK * 100) / 100)
}

func CelsiusToKelvin(value float64) float64 {
  CtoK := (value + 273.15)
  return (math.Round(CtoK * 100) / 100)
}

func CelsiusToFahrenheit(value float64) float64 {
  CtoF := (value * (9.0 / 5.0)) + 32
  return (math.Round(CtoF * 100) / 100)
}

func KelvinToFahrenheit(value float64) float64 {
  KtoF := (value - 273.15) * (9.0 / 5.0) + 32
  return (math.Round(KtoF * 100) / 100)
}

func KelvinToCelsius(value float64) float64 {
  KtoC := (value - 273.15)
  return (math.Round(KtoC * 100) / 100)
}

