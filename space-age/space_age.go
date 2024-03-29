// Package space contains the function to convert seconds to years on a planet
package space

type Planet string
const secondsOnThEarth =  31557600
var conversionPlanets = map[Planet]float64{
	"Earth":1,
	"Mercury":  0.2408467,
	"Venus": 0.61519726,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}
// Age convert seconds to years on a target planet
func Age(seconds float64, planet Planet)float64{
	return seconds /secondsOnThEarth / conversionPlanets[planet]
}