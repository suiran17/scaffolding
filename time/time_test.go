package time

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	
	s := time.Now().Format(Layout("Y-m-d H:i:s"))
	t.Log(s)
	
	ts := Date("Y-m-d", time.Now())
	t.Log(ts)
}
