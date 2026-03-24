package L2

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

/*
TODO: make sure the required fields have non-zero values
Return: 
	true: sender and recipient each contain a name and a number
	false: if any default zero values present
*/
func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 || 
	mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	} 
	return true
}
