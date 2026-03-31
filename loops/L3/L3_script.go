package L3

func countConnections(groupSize int) int {
	var connections int
	for groupSize > 0 {
		groupSize --
		connections += groupSize
	}
	return connections
}
