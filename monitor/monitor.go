package monitor

// monitors a fleet of coffee trees year round - temperature, humidity, sunlight amount, and soil quality

// create a coffee tree object
type CoffeeTree struct {
	// the age of the tree
	age int
	// the height of the tree, m
	height float64
	// the number of coffee beans on the tree
	numberOfCoffeeBeans int
	// soil quality
	soilQuality int
	// planting area, m2
	plantingSpace float64
	// plant health - scale of 1-10
	plantHealth int
}

// gets the volume taken up by the tree
func (tree CoffeeTree) getVolume() float64 {
	return tree.plantingSpace * tree.height
}

// calculates the soil quality after the tree is planted
func (tree CoffeeTree) calculateSoilQuality() int {
	// TODO: coffee plants need slightly acidic soil - this function should be calculated in real time from
	return tree.soilQuality - 1
}
