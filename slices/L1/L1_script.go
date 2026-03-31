package L4

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgs := [3]string{primary, secondary, tertiary}
	var costs [3]int
	curr_cost := 0

	for i:= 0; i < 3; i++ {
		curr_cost += len(msgs[i])
		costs[i] = curr_cost
	}
	return msgs, costs
}
