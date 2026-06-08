package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var totalBirds int
    for _, value := range birdsPerDay {
       totalBirds += value
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
   var result int
    start := (week -1) * 7
    week1 := week * 7
    for _, birds := range birdsPerDay[start:week1]{
       result += birds
    }
    return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
       birdsPerDay[i] ++
    }
    return birdsPerDay
}
