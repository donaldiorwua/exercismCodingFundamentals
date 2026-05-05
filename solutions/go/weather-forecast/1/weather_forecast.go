//Package weather provides tools that forcast weather conditions.
package weather

var (
    //CurrentCondition is the weather condition of that place at a given time.
	CurrentCondition string
    
    //CurrentLocation is a place where the weather conditions are being forcasted.
    CurrentLocation  string
)
//Forecast provides the weather conditions of a city in real time.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
