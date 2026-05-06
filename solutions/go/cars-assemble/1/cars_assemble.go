package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    total := float64(productionRate)
    result := total * successRate/100
    return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    total := float64(productionRate)
    result := total * successRate/100
    carspermin := result/60
    return int(carspermin)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    if carsCount > 10 {
		groupcars := carsCount / 10
		groupcarscost := groupcars * 95000
        reaminder := carsCount % 10
    	result := reaminder * 10000
    	total := groupcarscost + result
    	return uint(total)
    }else if carsCount == 10{
        carsCount = 95000
        return uint(carsCount)
    }else{
        carcost := carsCount * 10000
        return uint(carcost)
    }
    
}
