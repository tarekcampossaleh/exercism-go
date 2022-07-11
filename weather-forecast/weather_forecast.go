// Package weather package package package package package.
package weather

// CurrentCondition var var var var var var var var.
var CurrentCondition string

// CurrentLocation var var var var var var var var var.
var CurrentLocation string

// Forecast func func func func func func func.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
