package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    result := []int{}
    cards := []int{0,1,2,3,4,5,6,7,8,9}
	for i, val := range cards{
        if i == 2 || i == 6 || i == 9{
            result = append(result, val)
        }
    }    
    return result
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
        if index < 0 || index >= len(slice) {
            return -1
        }
    return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index >= len(slice) || index < 0 {
       return append(slice, value)
    }
    slice[index] = value
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    result := []int{}
    result = append(values, slice...)
    return result
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index < 0 || index > len(slice){
        return slice
    }
    return append(slice[:index], slice[index+1:]...)
}
