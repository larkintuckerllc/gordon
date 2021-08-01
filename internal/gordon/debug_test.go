package gordon

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	err := deleteRecords("training-main-317518", "instance-1", "10.0.0.1")
	if err != nil {
		fmt.Println(err)
		t.Error("err != nil")
		return
	}
}
