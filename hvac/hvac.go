package hvac

// the package hvac is responsible for heating and cooling the coffee trees

// Calculates the amount of energy needed to maintain a given temperature in a given volume of air, given the humidity and thermal efficiency of the container.
//
// volume: volume of air in cubic meters
// temp: temperature of air in degrees Celsius
// humidity: relative humidity of air as a percentage
// efficiency: thermal efficiency of the container as a percentage
//
// returns: amount of energy needed in joules
func HvacEnergyNeededPerHour(volume float64, targetTemp float64, idealTemp float64, humidity float64, efficiency float64) float64 {
	// Calculate the specific heat capacity of air at the given temperature and humidity
	specificHeatCapacity := CalculateSpecificHeatCapacity(targetTemp, humidity)
	// Calculate the amount of energy needed to maintain the temperature
	energy := volume * specificHeatCapacity * (targetTemp - idealTemp) / efficiency
	return energy
}

// depending on the temperature and humidity, the specific heat capacity of air changes
// this formula expects the temperature in degrees celsius and the humidity in percent
func CalculateSpecificHeatCapacity(temp float64, humidity float64) float64 {
	// Calculate the specific heat capacity of dry air at the given temperature
	dryAirHeatCapacity := 1005 + (temp-15)*0.34
	// Calculate the specific heat capacity of water vapor at the given temperature
	waterVaporHeatCapacity := 1925 + (temp-15)*0.34
	// Calculate the specific heat capacity of air at the given temperature and humidity
	specificHeatCapacity := dryAirHeatCapacity + (humidity/100)*waterVaporHeatCapacity
	return specificHeatCapacity
}
