/*
	Package gomessages sends messages on the goshimmer network
*/
package gomessages

import (
	"fmt"

	"github.com/iotaledger/goshimmer/client"
	"github.com/iotaledger/goshimmer/packages/tangle/payload"
)

var goshimAPI *client.GoShimmerAPI

func OpenChannel(theurl string) {
	goshimAPI = client.NewGoShimmerAPI(theurl)
}

func InfoVersion() string {

	info, err := goshimAPI.Info()
	if err != nil {
		// return error
	}

	// will print the response
	return "Node version " + info.Version + " on " + fmt.Sprintf("%d", info.NetworkVersion) + ": " + info.IdentityIDShort

}

func SendShimData(themessage string) string {
	thePayload := payload.NewGenericDataPayload([]byte(themessage))
	messageID, err := goshimAPI.SendPayload(thePayload.Bytes())
	errstr := ""
	if err != nil {
		// return error
		errstr = " - NOT!"
	}
	fmt.Println("Sent message " + messageID + errstr)
	return messageID
}

func SendShimMessage(themessage string) string {
	messageID, err := goshimAPI.Data([]byte(themessage))
	errstr := ""
	if err != nil {
		// return error
		errstr = " - NOT!"
	}
	fmt.Println("Sent message " + messageID + errstr)
	return messageID
}

func GetShimMessage(id string) string {
	message, err := goshimAPI.GetMessage(id)
	if err != nil {
		// return error
		return "Could not retrieve message"
	}
	return string(message.Payload)
}

// Test string function
func TestOutput() string {
	return "Hello there!"
}
