package constants

// pounds of coffee cherry per year in gram (10 lbs)
const PoundsOfCoffeeCherryPerYearPerTree = 4536

const CherryToGreenCoffeeWeightPercentage = 0.2

// RoastedCoffeeWeightPercentage is the percentage of roasted coffee weight from green to roasted
const RoastedToGreenCoffeeWeightPercentage = 0.8

// MinSellAmountPerDay is the minimum amount of coffee that should be sold per day
const MinSellAmountPerDay float64 = 1000

// MaxSellAmountPerDay is the maximum amount of coffee that should be sold per day
const MaxSellAmountPerDay float64 = 1814

// days per year lol
const DaysPerYear = 365

// average height of adult coffee tree in meters
const AverageHeightOfAdultCoffeeTree = 2

// average area of adult coffee tree in meters squared
const AverageAreaOfAdultCoffeeTree float64 = 1.0

// arabica is a highland coffee - thus it likes lower temperatures
type CoffeeGrowingConditions struct {
	MinTemperature float64
	MaxTemperature float64
	MinHumidity    float64
	MaxHumidity    float64
}

type MonthlyData struct {
	Month         string
	EnergyPerHour float64
	EnergyUsage   float64
	EnergyCost    float64
}

var ArabicaGrowingConditions CoffeeGrowingConditions = CoffeeGrowingConditions{
	MinTemperature: 18,
	MaxTemperature: 24,
	MinHumidity:    60,
	MaxHumidity:    80,
}

var RobustaGrowingConditions CoffeeGrowingConditions = CoffeeGrowingConditions{
	MinTemperature: 22,
	MaxTemperature: 30,
	MinHumidity:    70,
	MaxHumidity:    75,
}
