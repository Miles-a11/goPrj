package calculator

// CalculateInterest calculates the simple interest given principal, rate of interest and time period in years.
func CalculateInterest(principal float64, rateOfInterest float64, timeInYears int) float64 {
	return (principal * rateOfInterest * float64(timeInYears)) / 100
}