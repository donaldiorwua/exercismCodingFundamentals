package blackjack

//import "strings"


// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int
	switch card {
        case "ace":
        value = 11
        case "queen":
        value = 10
        case "jack":
        value = 10
        case "ten":
        value = 10
        case "king":
        value = 10
        case "two":
        value = 2
        case "three":
        value = 3
        case "four":
        value = 4
        case "five":
        value = 5
        case "six":
        value = 6
        case "seven":
        value = 7
        case "eight":
        value = 8
        case "nine":
        value = 9
        case "other":
        value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
   playerTotal := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
       switch {
        case card1 == "ace" && card2 == "ace":
           return "P"
        case  playerTotal == 21:
           if dealerValue == 10 || dealerValue == 11{
                return "S"
           }else{
            	return "W"
           }
        case playerTotal >= 17:
           return "S"
        case playerTotal >=12:
           if dealerValue >= 7{
               return "H"
           }
           return "S"
           default:
           return "H"
    }
}
