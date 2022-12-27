package main

func DecodeRoman(roman string) int {

	numbersMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize the result to 0
	result := 0

	// Loop through the characters in the Roman numeral string
	for i, r := range roman {
		// Get the Arabic numeral for the current character
		number := numbersMap[r]

		// If the next character has a higher value, subtract the current value from the result
		if i+1 < len(roman) && numbersMap[rune(roman[i+1])] > number {
			result -= number
		} else {
			// Otherwise, add the current value to the result
			result += number
		}
	}

	// Convert the result to a string and print it
	return result
}
