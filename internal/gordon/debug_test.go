package gordon

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	ip, err := getRecord("training-main-317518", "instance-1")
	if err != nil {
		fmt.Println(err)
		t.Error("err != nil")
		return
	}
	t.Error(*ip)
}
