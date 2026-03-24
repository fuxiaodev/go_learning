package L1

import (
	"fmt"
	"time"
)

// TODO: return the content of the message
// TODO: return the cost of the message, which is the length of the message multiplied by 3
func sendMessage(msg message) (string, int) {
	content := msg.getMessage()
	return content, 3 * len(content)
}

// TODO: declare this message interface
type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}
