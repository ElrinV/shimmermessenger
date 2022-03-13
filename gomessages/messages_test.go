package gomessages_test

import (
	"testing"

	"github.com/ElrinV/shimmermessenger/gomessages"
)

func Test_TestOutput(t *testing.T) {
	if gomessages.TestOutput() != "Hello there!" {
		t.Fatal("Wrong output")
	}
}
