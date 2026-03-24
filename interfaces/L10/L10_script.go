package L10

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

// TODO: implement the importance methods for each message type. 
// they should return the importance score for each message type
func (d directMessage) importance() int {
	if d.isUrgent{
		return 50
	}
	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

// TODO: complete the processNotification function.
// It should identify the type and return different values for each type
func processNotification(n notification) (string, int) {
	switch t := n.(type) {
	case directMessage:
			return t.senderUsername, t.importance()
	case groupMessage:
		return t.groupName, t.importance()
	case systemAlert:
		return t.alertCode, t.importance()
	default:
		return "", 0
	}
}

// interface.(type) is a special go syntax for type switch
// It inspects the runtime type of n (which is an interface).
// t:== declares a new var that will hold n cast to the matched concrete type inside the case
// It cannot be used to assign a concrete type to a variable outside a type switch
/*
outside a type switch use:
if dm, ok := n.(directMessage); ok {
    // dm is of type directMessage
}
*/