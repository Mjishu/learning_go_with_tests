package clock

import (
	"testing"
	"time"

	"github.com/mjishu/intro_learn_go_with_tests/maths/clockface"
)

func TestClock(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
