package L2

/*
Loops in go can omit sections of a for loop.
for example, the condition can be omitted which causes the loop to run forever
*/

func maxMessages(thresh int) int {
	var total int
	var count int
	// ever ending loop
	for i:= 0; ; i ++{
		if total + 100 + i <= thresh {
			total += 100 + i
			count += 1
		}else{
			return count
		}
	} 
}
