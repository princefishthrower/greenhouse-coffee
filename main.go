package main

import (
	"coffee/charting"
	"coffee/constants"
	"coffee/costs"
	"coffee/hvac"
	"coffee/sales"
	"coffee/weather"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// show the heating costs per month
	coffeeType := "arabica" // try robusta - may be cheaper for you if you live in a warmer climate

	// targetYearlyCropInGrams := constants.MaxSellAmountPerDay * constants.DaysPerYear
	// targetYearlyCropInGrams := constants.MinSellAmountPerDay * constants.DaysPerYear
	targetYearlyCropInGrams := 1000.0 * constants.DaysPerYear
	greenhouseThermalEfficiency := 0.8
	salePriceRoastedBeansPerKg := costs.SalePricePerKilogramTirolMountainCoffeeEUR

	displayHeatingCostsPerMonth(coffeeType, targetYearlyCropInGrams, greenhouseThermalEfficiency, salePriceRoastedBeansPerKg)
}

func displayHeatingCostsPerMonth(coffeeType string, targetYearlyCropGram float64, greenhouseThermalEfficiency float64, salePriceRoastedBeansPerKg float64) {
	// print in caps the coffee type
	fmt.Printf("Coffee type: %s\n\n", strings.ToUpper(coffeeType))
	fmt.Printf("Target Crop: %f g (%f kg)\n\n", targetYearlyCropGram, targetYearlyCropGram/1000)

	targetGramsOfGreenCoffeeBeans := targetYearlyCropGram / constants.RoastedToGreenCoffeeWeightPercentage
	targetGramsOfCoffeeCherry := targetGramsOfGreenCoffeeBeans / constants.CherryToGreenCoffeeWeightPercentage

	// print each of these numbers out
	fmt.Printf("Target amount finished coffee to sell: %f (%f kg) \n", targetYearlyCropGram, targetYearlyCropGram/1000)
	fmt.Printf("Target amount green coffee beans: %f (%f kg)\n", targetGramsOfGreenCoffeeBeans, targetGramsOfGreenCoffeeBeans/1000)
	fmt.Printf("Target amount coffee cherry: %f (%f kg)\n", targetGramsOfCoffeeCherry, targetGramsOfCoffeeCherry/1000)

	numTrees := targetGramsOfCoffeeCherry / constants.PoundsOfCoffeeCherryPerYearPerTree

	// round up to the nearest tree
	if numTrees-float64(int(numTrees)) > 0 {
		numTrees = float64(int(numTrees)) + 1
	}

	// log out the number of trees needed
	fmt.Printf("Number of trees needed: %f\n\n", numTrees)

	totalVolumeOfGreenHouse := numTrees * constants.AverageAreaOfAdultCoffeeTree * constants.AverageHeightOfAdultCoffeeTree

	// log out total volume of greenhouse
	fmt.Printf("Total volume of greenhouse: %f m^3\n\n", totalVolumeOfGreenHouse)

	// select the growing conditions based on the coffee type
	growingConditions := constants.ArabicaGrowingConditions
	if coffeeType == "robusta" {
		growingConditions = constants.RobustaGrowingConditions
	}

	averageTemperatureToMaintain := (growingConditions.MinTemperature + growingConditions.MaxTemperature) / 2
	averageHumidityToMaintain := (growingConditions.MinHumidity + growingConditions.MaxHumidity) / 2

	// for each month, use the weather package data combined with the hvac package to determine the energy needed to maintain the temperature each month
	// also keep track of total energy and cost
	totalEnergy := 0.0
	totalCostEUR := 0.0

	// initialize monthlyData array
	monthlyData := make([]constants.MonthlyData, len(weather.MonthlyWeatherData))

	for i := 0; i < len(weather.MonthlyWeatherData); i++ {
		weather := weather.MonthlyWeatherData[i]
		energyNeededPerHourInJoules := hvac.HvacEnergyNeededPerHour(totalVolumeOfGreenHouse, averageTemperatureToMaintain, (weather.Low+weather.High)/2, averageHumidityToMaintain, greenhouseThermalEfficiency)
		energyNeededPerHourInKiloWatts := energyNeededPerHourInJoules / 3600000
		energyNeededPerMonthInJoules := energyNeededPerHourInJoules * 24 * 30
		// convert to kiloWattHours
		energyNeededPerMonthInKiloWattHours := energyNeededPerMonthInJoules / 3600000
		// calculate the cost of the energy - from most recent data, 25.3 euro per 100 kWh
		energyCostPerMonth := energyNeededPerMonthInKiloWattHours * (25.3 / 100)

		// add to total energy and cost
		totalEnergy += energyNeededPerMonthInKiloWattHours
		totalCostEUR += energyCostPerMonth

		// add to monthlyData array
		monthlyData[i] = constants.MonthlyData{
			Month:         weather.Month,
			EnergyPerHour: energyNeededPerHourInKiloWatts,
			EnergyUsage:   energyNeededPerMonthInKiloWattHours,
			EnergyCost:    energyCostPerMonth,
		}

		// print out the results
		fmt.Printf("%s - HVAC Energy Needed Per Month: %f kWh, Cost: %f EUR, Per hour: %f kW \n", weather.Month, energyNeededPerMonthInKiloWattHours, energyCostPerMonth, energyNeededPerHourInKiloWatts)
	}

	// print out total energy needed
	totalEnergyKWh := totalEnergy / 1000
	fmt.Printf("\nTotal HVAC Energy Needed: %f kWh (%f MWh)\n", totalEnergy, totalEnergyKWh)

	// print out the total cost
	fmt.Printf("Total Energy Cost: %f EUR\n\n", totalCostEUR)

	// print out breakeven cost per 100 grams of coffee
	// breakevenCostPer100GramsOfCoffee := totalCostEUR / (targetYearlyCropGram / 100)
	// breakevenCostPer500GramsOfCoffee := breakevenCostPer100GramsOfCoffee * 5
	// breakevenCostPerKgOfCoffee := breakevenCostPer500GramsOfCoffee * 2
	// fmt.Printf("Breakeven Costs:\n100g @ %f EUR\n500g @ %f EUR\n1kg @ %f EUR\n\n", breakevenCostPer100GramsOfCoffee, breakevenCostPer500GramsOfCoffee, breakevenCostPerKgOfCoffee)
	// fmt.Printf("Compare to Coconut Crunch Coffee Pricing:\n100g @ $4.85\n500g @ $24.30\n1kg @ $48.59\n\n")
	// fmt.Printf("You're gonna be rich!\n\n")

	// print profit per year
	profitPerYear := (sales.GramsCoffeeSoldPerDay / 1000) * 365 * salePriceRoastedBeansPerKg
	fmt.Printf("Profit per year: %f EUR\n\n", profitPerYear)

	// find startup costs based on kw needed - first get the max of EnergyPerHour of the months array
	maxEnergyPerHour := 0.0
	maxEnergyMonth := ""
	for _, month := range monthlyData {
		if month.EnergyPerHour > maxEnergyPerHour {
			maxEnergyPerHour = month.EnergyPerHour
			maxEnergyMonth = month.Month
		}
	}

	// get the startup cost based on the max energy per hour
	solarSystemSizeKW, startupCost := costs.FindStartupCost(maxEnergyPerHour)

	// print out the month with the highest energy per hour, what the max energy per hour is, and the startup cost
	fmt.Printf("Month found with highest energy per hour usage: %s\n", maxEnergyMonth)
	fmt.Printf("Max energy per hour: %f kW\n", maxEnergyPerHour)
	fmt.Printf("Startup Cost (for a %f kW solar system): %f EUR\n\n", solarSystemSizeKW, startupCost)

	// number of days to breakeven is the startup cost minus the amount of coffee sold times the coconut crunch example price
	daysToBreakeven := startupCost / ((sales.GramsCoffeeSoldPerDay / 1000) * salePriceRoastedBeansPerKg)

	// print out the number of days to breakeven
	fmt.Printf("Days to breakeven: %f (%f years)\n\n", daysToBreakeven, daysToBreakeven/365)

	// print summary - initial upfront costs and then the profits that will appear after the breakeven point
	fmt.Printf("Summary:\n")
	fmt.Printf("Startup Cost (for a 20 kW solar system): %f EUR\n", startupCost)
	fmt.Printf("Profit per year after break even: %f EUR\n", profitPerYear)
	fmt.Printf("Days to breakeven: %f (%f years)\n\n", daysToBreakeven, daysToBreakeven/365)

	// create and open monthly energy cost chart
	charting.BuildMonthlyEnergyCostChart(monthlyData)
	cmd := exec.Command("open", "monthly_energy_cost.html")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create and open monthly energy usage chart
	charting.BuildMonthlyEnergyUsageChart(monthlyData)
	cmd = exec.Command("open", "monthly_energy_usage.html")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create and open totals chart
	charting.BuildTotalsChart(totalEnergyKWh, totalCostEUR)
	cmd = exec.Command("open", "totals.html")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// here we go with thermodynamics, calculate the energy needed to maintain the greenhouse volume at the average temperature
	// 1 m3 of air at 20 degrees celsius has 1.005 kJ of energy
	// 1 kJ of energy is 0.23900573614 BTU
	// 1 BTU is 0.29307107 kWh
	// 1 kWh is 0.00027777778 MWh
}
