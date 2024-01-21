package logging

import "fmt"

func LogMe(msg, msgtype string) {
	fmt.Printf("%s: %s\n",msgtype, msg)
}
