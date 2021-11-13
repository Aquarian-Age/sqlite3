package person

import (
	"fmt"
	"testing"
)

func TestDBdo(t *testing.T) {
	//var t1 time.Time = time.Date(2021, time.Month(9), 3, 0, 0, 0, 0, time.Local)
	objMap := NewPersonDB()

	fmt.Println("==========================")
	for k, v := range objMap {
		fmt.Println(k, v)
	}
}
