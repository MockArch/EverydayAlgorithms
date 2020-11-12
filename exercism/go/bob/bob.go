package bob

func Hey(remark string) string {

	remarkByte := []byte(remark)
	return checkBobResp(remarkByte)
}

func checkBobResp(s []byte) string {
	var msgbob string
	lenSlice := len(s) - 1
	isFullCap := true
	for d, i := range s {
		if lenSlice != d {
			if i >= 65 && i <= 90 && isFullCap {
				msgbob = "Whoa, chill out!"
			} else {
				isFullCap = false
				msgbob = "Whatever."
			}
		} else {
			if i == 63 {
				msgbob = "Sure."
			}

		}
	}
	return msgbob
}
