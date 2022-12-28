package costs

// map of kw solar system to initial cost in EUR
// url pattern: https://quotes.solarproof.com.au/system-sizes/50kw-solar-system-information-facts-figures/
// just change the 50 to the kW you want
var StartupCosts = map[float64]float64{
	1:  1138.0,
	5:  5567.0,
	10: 11071.0,
	20: 22176.0,
	30: 33214.0,
	40: 44289.0,
	50: 55362.0,
}

const SalePricePerGramOfCoffeeCoconutCrunchEUR float64 = 0.046
const SalePricePerKilogramOfCoffeeCoconutCrunchEUR float64 = 0.046 * 1000
const SalePricePerKilogramJamaicaBlueMountainCoffeeEUR float64 = 108.00
const SalePricePerKilogramTirolMountainCoffeeEUR float64 = 54.00
const SalePricePerKilogramGlobalAverage float64 = 4.0

func FindStartupCost(targetKW float64) (float64, float64) {

	nearestMatch := 0.0

	for k := range StartupCosts {
		if k >= targetKW && (k < nearestMatch || nearestMatch == 0) {
			nearestMatch = k
		}
	}

	return nearestMatch, StartupCosts[nearestMatch]
}
