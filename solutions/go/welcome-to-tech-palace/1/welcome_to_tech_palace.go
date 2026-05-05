package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    mgs := "Welcome to the Tech Palace, "
   var name string
	name = customer
    return mgs + strings.ToUpper(name)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    
    result :=  strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + 		     strings.Repeat("*", numStarsPerLine)
    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    result := strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
    return result
}
