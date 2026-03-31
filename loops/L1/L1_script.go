package L1

func bulkSend(numMessages int) float64 {
	var total float64
	for i := 0; i < numMessages; i ++ {
		total += 1.0 + float64(i)/100 
	}
	return total
}