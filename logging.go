package logging

import "fmt"

func LogMe(msg, msgtype string) {
	fmt.Println(msgtype, ":", msg)
}
