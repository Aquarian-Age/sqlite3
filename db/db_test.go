package db

import (
	"testing"
	"time"
)

func TestDBdo(t *testing.T) {
	var t1 time.Time = time.Date(2021, time.Month(9), 3, 0, 0, 0, 0, time.Local)
	DBdo(t1)
}
