package math

// Avg comment ...
func Avg(xs []float64) float64 {
	ttl := float64(0)
	for _, x := range xs {
		ttl += x
	}
	return ttl / float64(len(xs))
}
