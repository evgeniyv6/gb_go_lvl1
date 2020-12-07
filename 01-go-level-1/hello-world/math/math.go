package math

// Avg test
func Avg(nums []float64) float64 {
	sum := float64(0)
	for _, el := range nums {
		sum += el
	}
	return sum / float64(len(nums))

}
