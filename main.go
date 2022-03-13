package main

import (
	"fmt"

	"github.com/ElrinV/shimmermessenger/gomessages"
)

func main() {
	gomessages.OpenChannel("http://192.168.2.120:8080")

	fmt.Println(gomessages.InfoVersion())

	//gomessages.SendShimMessage("Hello Shimmer World")
	msgID := gomessages.SendShimData("Hello Shimmer World")

	fmt.Println(gomessages.GetShimMessage(msgID))

	msgID = gomessages.SendShimMessage("Hallo Shimmer World")

	fmt.Println(gomessages.GetShimMessage(msgID))

}
