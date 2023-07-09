package controller

import "fmt"

func joinMsg(args ...interface{}) string {
	msg := ""
	for i, a := range args {
		if i == 0 {
			msg = fmt.Sprint(a)
		} else {
			msg = fmt.Sprintf("%v %v", msg, a)
		}
	}
	return msg
}
