package L4

func getMessageCosts(messages []string) []float64 {
	message_costs := make([]float64, len(messages))
	
	for index, item := range messages {
		message_costs[index] = float64(len(item)) * 0.01
	}

	return message_costs
}
/*
len and cap are legal when applied to the nil slice, and return 0
*/