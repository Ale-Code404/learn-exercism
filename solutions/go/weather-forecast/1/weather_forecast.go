// Package weather has tools to format weather conditions to a human readable sentences.
package weather

var (
    // CurrentCondition has the current weather condition.
	CurrentCondition string
    // CurrentLocation has the current location on earth.
	CurrentLocation  string
)

// Forecast returns a formatted sentence about the  weather thats includes a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
